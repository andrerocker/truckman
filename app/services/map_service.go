package services

import (
	"../models"
	"../repositories"
	"bufio"
	"strconv"
	"strings"
)

type MapService struct {
	Repo repository.MapRepository
}

func NewMapService(name string) (MapService, error) {
	repository, err := repository.NewMapRepository(name)

	if err != nil {
		return MapService{}, err
	}

	return MapService{repository}, nil
}

func (self MapService) InitializeDatabase(mapContent string) {
	scanner := bufio.NewScanner(strings.NewReader(mapContent))
	for scanner.Scan() {
		nodes := strings.Split(scanner.Text(), " ")
		value, _ := strconv.Atoi(nodes[2])

		self.Repo.CreateAndRelate(nodes[0], nodes[1], value)
	}
}

func (self MapService) TraceBetterRoute(routeForm models.RouteForm) ([]models.Route, error) {
	routes, err := self.Repo.TraceBetterRoute(routeForm)

	if err != nil {
		return routes, err
	}

	for index, route := range routes {
		// http://stackoverflow.com/questions/15945030/change-values-while-iterating-in-golang
		routes[index].Cost = route.Distance * routeForm.FuelCost / routeForm.Autonomy
		routes[index].Path = strings.TrimSpace(route.Path)
	}

	return routes, nil
}
