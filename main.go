package main
import(
	"mongodb_assign/mongodbF5interns"
)
func main() {

	db := new(mongodb)
	//temp(&ash)
	connector(db)
	insertor(db)
	updator(db)
	retreivor(db)
	deletor(db)
}

/*func temp (arg1 interface { }){
    arg2 := (Trainer*)arg1
    fmt.Printf("%v", arg2)
}*/

func connector(n connection) {
	n.connect()
}
func insertor(n insertion) {
	ash := Trainer{"Ash", 10, "Pallet Town"}
	n.insert(ash)
}
func updator(n updation) {
	n.update()
}
func retreivor(n retreival) {
	n.retreive()
}
func deletor(n deletion) {
	n.delete()
}
