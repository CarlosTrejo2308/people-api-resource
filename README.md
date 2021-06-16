# peopleApiResource

This is a helper package designed to be used on this docker ejercise: https://github.com/AnhellO/dockercises/tree/development/docker-compose/ejercicio-1

It includes 3 modules:
1. abort: To handle errors, currently it aborts when it founds one
2. DB: It gets the mongo uri string from eviroment values and gets a mongo connection 
3. People: Defines structs to be used in the xml file
