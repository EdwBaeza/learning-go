package maps

// obtiene un map
func GetMap() map[string]int {

	// es necesario inicializar
	var mapTest = make(map[string]int)
	mapTest["llave1"] = 1
	mapTest["llave2"] = 200
	// otherMap := map[string]int{
	// 	"Juan": 18,
	// }
	// fmt.Println(otherMap)
	delete(mapTest, "llave2")
	return mapTest
}
