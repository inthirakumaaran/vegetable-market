package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
)

import . "vegetable-market/common"

type MARKET int

var market []Vegetable
var lock sync.Mutex

var Marshal = func(v []Vegetable) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

var Unmarshal = func(r io.Reader, v []Vegetable) error {
	return json.NewDecoder(r).Decode(&v)
}

func (a *MARKET) GetAllVegetablesName(placeholder string, reply *[]string) error {
	err := Load(MARKET_DB_PATH, market)
	if err != nil {
		return err
	}
	var vegList []string

	for _, val := range market {
		vegList = append(vegList, val.Name)
	}

	*reply = vegList
	return nil
}

func (a *MARKET) GetAllVegetables(placeholder string, reply *[]Vegetable) error {
	err := Load(MARKET_DB_PATH, market)
	if err != nil {
		return err
	}
	var vegList []Vegetable

	for _, val := range market {
		vegList = append(vegList, val)
	}

	*reply = vegList
	return nil
}

func (a *MARKET) GetVegetableDetails(veg string, reply *Vegetable) error {
	err := Load(MARKET_DB_PATH, market)
	if err != nil {
		return err
	}
	var selectedVeg Vegetable
	var vegFound bool = false
	for idx, val := range market {
		if (val.Name == veg) {
			vegFound = true
			selectedVeg = market[idx]
		}
	}
	if !vegFound {
		return errors.New("Vegetable " + veg + " is not found")
	}
	*reply = selectedVeg
	return nil
}

func (a *MARKET) NewVegetable(veg Vegetable, reply *Vegetable) error {
	err := Load(MARKET_DB_PATH, market)
	if err != nil {
		return err
	}
	var vegFound bool = false
	for _, val := range market {
		if (val.Name == veg.Name) {
			vegFound = true
		}
	}
	if vegFound {
		return errors.New("Vegetable " + veg.Name + " already exists")
	}
	market = append(market, veg)
	err1 := Save(MARKET_DB_PATH, market)
	if err1 != nil {
		return err1
	}
	*reply = veg
	return nil
}

func (a *MARKET) UpdateVegetable(veg Vegetable, reply *Vegetable) error {
	err := Load(MARKET_DB_PATH, market)
	if err != nil {
		return err
	}
	var editVeg Vegetable
	var vegFound bool = false
	for idx, val := range market {
		if (val.Name == veg.Name) {
			market[idx] = Vegetable{veg.Name, veg.Price, veg.Quantity}
			vegFound = true
			editVeg = market[idx]
			err := Save(MARKET_DB_PATH, market)
			if err != nil {
				return err
			}
		}
	}
	if !vegFound {
		return errors.New("Vegetable " + veg.Name + " is not found")
	}
	*reply = editVeg
	return nil
}

func (a *MARKET) GetVegetablePrice(veg string, reply *int) error {
	err := Load(MARKET_DB_PATH, market)
	if err != nil {
		return err
	}
	var price int
	var vegFound bool = false
	for _, val := range market {
		if (val.Name == veg) {
			price = val.Price
			vegFound = true
		}
	}
	if !vegFound {
		return errors.New("Vegetable " + veg + "  is not found")
	}
	*reply = price
	return nil
}

func (a *MARKET) GetVegetableQuantity(veg string, reply *int) error {
	err := Load(MARKET_DB_PATH, market)
	if err != nil {
		return err
	}
	var quantity int
	var vegFound bool = false
	for _, val := range market {
		if (val.Name == veg) {
			quantity = val.Quantity
			vegFound = true
		}
	}
	if !vegFound {
		return errors.New("Vegetable " + veg + " is not found")
	}
	*reply = quantity
	return nil
}


func (a *MARKET) UpdateVegetableQuantity(veg UpdateVegetable, reply *Vegetable) error {
	err := Load(MARKET_DB_PATH, market)
	if err != nil {
		return err
	}
	var editVeg Vegetable
	var vegFound bool = false
	for idx, val := range market {
		if (val.Name == veg.Name) {
			market[idx] = Vegetable{veg.Name, val.Price, veg.Value}
			vegFound = true
			editVeg = market[idx]
			err := Save(MARKET_DB_PATH, market)
			if err != nil {
				return err
			}
		}
	}
	if !vegFound {
		return errors.New("Vegetable " + veg.Name + " is not found")
	}
	*reply = editVeg
	return nil
}

func (a *MARKET) UpdateVegetablePrice(veg UpdateVegetable, reply *Vegetable) error {
	err := Load(MARKET_DB_PATH, market)
	if err != nil {
		return err
	}
	var editVeg Vegetable
	var vegFound bool = false
	for idx, val := range market {
		if (val.Name == veg.Name) {
			market[idx] = Vegetable{veg.Name, veg.Value, val.Quantity}
			vegFound = true
			editVeg = market[idx]
			err := Save(MARKET_DB_PATH, market)
			if err != nil {
				return err
			}
		}
	}
	if !vegFound {
		return errors.New("Vegetable " + veg.Name + " is not found")
	}
	*reply = editVeg
	return nil
}

func Save(path string, v []Vegetable) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

func Load(path string, v []Vegetable) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return Unmarshal(f, v)
}

func Delete(path string) error {
	lock.Lock()
	defer lock.Unlock()
	err := os.Remove(path)
	return err
}

func main() {
	Delete(MARKET_DB_PATH)
	Save(MARKET_DB_PATH, market)
	market := new(MARKET)
	err := rpc.Register(market)
	if err != nil {
		log.Fatal("error registering Market", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("Vegitable market is available on port %d", 9090)
	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error opening Market: ", err)
	}
}
