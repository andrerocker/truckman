## Passo a passo para realizar os testes

1 - Configure um servidor neo4j local sem autenticação
```
	- O problema encaminhado pode ser resolvido com um banco de dados de grafos
	- O melhor algoritmo possivel pra resolver esse problema é o Djisktra
	- http://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
```

2 - Inicialize o processo com um simples `make run` 
```
	- Ao realizar esse processo as dependencias do projeto seram baixadas
	- Realizei a implementação do webservice utilizando Go
	- O projeto necessita do go ~> 1.4 (a forma mais simples de instalar é com o GVM)
	- O serviço possui dois endpoints, um para carga e outro pra consulta
```

3 - No diretorio `util` existe dois scripts:
```
	- criar esses scripts foi a forma didatica que encontrei para ensinar a consumir o serviço
	- `populate`: script exemplo pra carregar uma malha
	- `statroute`: traça a melhor rota pros parametros solicitados 
```
