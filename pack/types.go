package pack


type Packer interface {
	init()	error
	Pack()	error
	UnPack()	error
}
