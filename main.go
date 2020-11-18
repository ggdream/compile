package main


func main() {
	if err := runCli(); err != nil {
		println(err.Error())
	}
}
