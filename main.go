package main

import "fmt"

/*
button içerisine uygulama(execute) fonksiyonunu çağıracak olan
Command interface' ini ekliyoruz.
*/
type Button struct {
	command Command
}

type Command interface {
	execute()
}

// press ile execute() fonksiyonunu tetikliyoruz.
func (b *Button) press() {
	b.command.execute()
}

/*
Device interface ile execute içerisinde işleme alınacak
fonksiyonları tanımlıyoruz.
*/
type Device interface {
	on()
	off()
}

// OnCommand struct ı ile Device interface içerisinde On() fonksiyonuna erişiyoruz.
type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

// OnCommand struct ı ile Device interface içerisinde Off() fonksiyonuna erişiyoruz.

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

// Gate için kontrol boolen değer oluşturuyoruz.
type Gate struct {
	isRunning bool
}

/*
Device interface içerisinde tanımladığımız On() fonksiyonun
Gate üzerindeki bool true işlemini gerçekleştiriyoruz.
*/
func (t *Gate) on() {
	t.isRunning = true
	fmt.Println("on")
}

/*
Device interface içerisinde tanımladığımız Off() fonksiyonun
Gate üzerindeki bool false işlemini gerçekleştiriyoruz.
*/
func (t *Gate) off() {
	t.isRunning = false
	fmt.Println("off")
}

func main() {
	gate := &Gate{}

	// OnCommand içerisine button true işlemi için gate struct ekliyoruz.
	onCommand := &OnCommand{
		device: gate,
	}

	// OffCommand içerisine button false işlemi için gate struct ekliyoruz.
	offCommand := &OffCommand{
		device: gate,
	}

	// Button struct içerisine execute işlemlerinden On() fonksiyonu çağıran OnCommand ekliyoruz.
	onButton := &Button{
		command: onCommand,
	}

	onButton.press() // Run On()

	// Button struct içerisine execute işlemlerinden Off() fonksiyonu çağıran OnCommand ekliyoruz.
	offButton := &Button{
		command: offCommand,
	}

	offButton.press() // Run Off()

}
