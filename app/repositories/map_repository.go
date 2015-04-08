package repository

import (
	"../models"
	"github.com/jmcvetta/neoism"
	"os"
	"strings"
)

type MapRepository struct {
	NameSpace string
	Database  *neoism.Database
}

func NewMapRepository(namespace string) (MapRepository, error) {
	databaseUrl := os.Getenv("NEO4J_URL")
	database, err := neoism.Connect(databaseUrl)

	if err != nil {
		return MapRepository{}, err
	}

	return MapRepository{namespace, database}, nil
}

func (self MapRepository) CreateAndRelate(origin, dest string, distance int) {
	node := self.createNode(origin)
	nody := self.createNode(dest)

	self.createRelation(node, nody, distance)
}

func (self MapRepository) TraceBetterRoute(routeForm models.RouteForm) ([]models.Route, error) {
	response := []models.Route{}

	query := `MATCH path = (start:{mapName}{name:"{origin}"})-[*]->(end:{mapName} {name:"{target}"})
		 RETURN
			reduce(names='', node in nodes(path) | names + node.name + ' ') AS path,
			reduce(sum=0, targets in relationships(path) | sum+targets.value) AS distance
		 ORDER BY distance ASC
		 LIMIT 1`

	// We have a bug on Neo4J rest api :(
	adapted := strings.Replace(query, "{mapName}", self.NameSpace, -1)
	adapted = strings.Replace(adapted, "{origin}", routeForm.Origin, -1)
	adapted = strings.Replace(adapted, "{target}", routeForm.Target, -1)

	cypherQuery := neoism.CypherQuery{
		Statement: adapted,
		Result:    &response,
	}

	err := self.Database.Cypher(&cypherQuery)

	if err != nil {
		return make([]models.Route, 0), err
	}

	return response, nil
}

func (self MapRepository) createNode(name string) *neoism.Node {
	node, _, _ := self.Database.GetOrCreateNode(name, "name", neoism.Props{"name": name})
	node.AddLabel(self.NameSpace)

	return node
}

func (self MapRepository) createRelation(node, nody *neoism.Node, distance int) {
	rels, _ := node.Relationships("target")

	for _, current := range rels {
		relation, _ := current.End()

		if nody.Id() == relation.Id() {
			return
		}
	}

	node.Relate("target", nody.Id(), neoism.Props{"value": distance})
}
