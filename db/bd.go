package bd
var datos map[string]string

func Init()  {
	 datos = make(map[string]int)
}
func Create(key string, data int) (string) {
	
	datos[key] = data
	
return "ingresado"
}

func Retrieve( key string)  int{
	
	dato := datos[key]
	
	return dato
}

func Update(key string, data int) string {
	datos[key] = data	
return "ingresado"
}

func Delete(key string)  {
	delete(datos, key)
}



