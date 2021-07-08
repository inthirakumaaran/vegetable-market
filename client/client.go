package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
	"github.com/urfave/cli/v2"
	"net/rpc"
	"vegetable-market/common"
)

func main() {

	app := cli.NewApp()
	app.Name = "Market Lookups"
	app.Usage = "Let's query vegetable and prices & quantity."
	app.Description = "Cmd app for our online market."
	app.Authors = []*cli.Author{
		{Name: "inthirakumaaran", Email: "tharmakulasingham.21@cse.mrt.ac.lk "},
	}

	app.Action = mainAction

	app.Commands = []*cli.Command{
		getCommand(),
		addCmd(),
		updateCmd(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getCommand() *cli.Command {
	return &cli.Command{
		Name: "get",
		Aliases: []string{"g"},
		Action: func(ctx *cli.Context) error {
			n := ctx.NArg()
			if n == 1 {
				a := ctx.Args().Get(0)
				var replyVeg common.Vegetable
				client, err := rpc.DialHTTP("tcp", "localhost:9090")

				if err != nil {
					log.Fatal("Connection error: ", err)
				}
				erro := client.Call("MARKET.GetVegetableDetails", a, &replyVeg)
				if erro != nil {
				log.Fatal("Error: ", erro)
				}  else {
				fmt.Println("Vegatable " + a + " details:", replyVeg)
				}

				return nil
			} else if n == 0 {
				var replyList []common.Vegetable
				client, err := rpc.DialHTTP("tcp", "localhost:9090")

				if err != nil {
					log.Fatal("Connection error: ", err)
				}

				erro := client.Call("MARKET.GetAllVegetables", "", &replyList)
				if erro != nil {
					log.Fatal("Error: ", erro)
				}  else {
					fmt.Println("Vegetables in Market: ", replyList)
				}
				
				return nil
			} else {
				return fmt.Errorf("Too many arguments provided for the get operation")
			}
		},
		Subcommands: []*cli.Command {
			{
				Name:  "names",
				Usage: "Get names of the all the available vegitables",
				Action: func(ctx *cli.Context) error {
					var replyList []string
					client, err := rpc.DialHTTP("tcp", "localhost:9090")

					if err != nil {
						log.Fatal("Connection error: ", err)
					}
                    erro := client.Call("MARKET.GetAllVegetablesName", "", &replyList)
					if erro != nil {
						log.Fatal("Error: ", erro)
					}  else {
						fmt.Println("Vegetables in Market: ", replyList)
					}
					
					return nil
				  },
			  },
			{
			  Name:  "quantity",
			  Usage: "Get quantity of the vegitable",
			  Action: func(ctx *cli.Context) error {
					n := ctx.NArg()
					if n != 1 {
						return fmt.Errorf("Required arguments are not provided for the get operation")
					}

					a := ctx.Args().Get(0)

					var replyQuantity int
					client, err := rpc.DialHTTP("tcp", "localhost:9090")

					if err != nil {
						log.Fatal("Connection error: ", err)
					}

					erro := client.Call("MARKET.GetVegetableQuantity", a, &replyQuantity)
					if erro != nil {
					log.Fatal("Error: ", erro)
					}  else {
					fmt.Println("Quantity of " + a + ":", replyQuantity)
					}
					
					return nil
				},
			},
			{
				Name:  "price",
				Usage: "Get price of the vegitable",
				Action: func(ctx *cli.Context) error {
					  n := ctx.NArg()
					  if n != 1 {
						  return fmt.Errorf("Required arguments are not provided for the get operation")
					  }

					  a := ctx.Args().Get(0)

					  var replyPrice int
					  client, err := rpc.DialHTTP("tcp", "localhost:9090")

					  if err != nil {
						  log.Fatal("Connection error: ", err)
					  }
					  erro := client.Call("MARKET.GetVegetablePrice", a, &replyPrice)
					  if erro != nil {
						log.Fatal("Error: ", erro)
					  }  else {
						fmt.Println("Price of " + a + ":", replyPrice)
					  }

					  return nil
			  },
			},
		},
	}
}

func addCmd() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Action: func(ctx *cli.Context) error {

			n := ctx.NArg()
			if n != 3 {
				return fmt.Errorf("Required arguments are not provided for add operation")
			}

			b := ctx.Args().Get(0)
			c := ctx.Args().Get(1)
			d := ctx.Args().Get(2)

			price, _ := strconv.Atoi(c)
			quantity, _ := strconv.Atoi(d)

			a := common.Vegetable{b, price, quantity}

			var replyVeg common.Vegetable
			client, err := rpc.DialHTTP("tcp", "localhost:9090")

			if err != nil {
				log.Fatal("Connection error: ", err)
			}

			erro := client.Call("MARKET.NewVegetable", a, &replyVeg);
			if erro != nil {
				log.Fatal("Error: ", erro)
			}  else {
				fmt.Println("Added vegetable: ", replyVeg)
			}

			return nil
		},
	}
}

func updateCmd() *cli.Command {

	return &cli.Command{
		Name:    "update",
		Aliases: []string{"u"},
		Action: func(ctx *cli.Context) error {

			n := ctx.NArg()
			if n != 3 {
				return fmt.Errorf("Required arguments are not provided for update operation")
			}

			b := ctx.Args().Get(0)
			c := ctx.Args().Get(1)
			d := ctx.Args().Get(2)

			price, _ := strconv.Atoi(c)
			quantity, _ := strconv.Atoi(d)

			a := common.Vegetable{b, price, quantity}

			var replyVeg common.Vegetable
			var replyPrice int
			var replyQuantity int
			client, err := rpc.DialHTTP("tcp", "localhost:9090")

			if err != nil {
				log.Fatal("Connection error: ", err)
			}

			erro := client.Call("MARKET.GetVegetablePrice", b, &replyPrice)
			if erro != nil {
				log.Fatal("Error: ", erro)
			}  else {
				fmt.Println("Old Price of " + b + ":", replyPrice)
			}

			erro1 := client.Call("MARKET.GetVegetableQuantity", b, &replyQuantity)
			if erro1 != nil {
				log.Fatal("Error: ", erro1)
			}  else {
				fmt.Println("Old quantity of " + b + ":", replyQuantity)
			}

			erro2 := client.Call("MARKET.UpdateVegetable", a, &replyVeg)
			if erro2 != nil {
				log.Fatal("Error: ", erro2)
			}  else {
				fmt.Println("New Price of " + b + ":", replyVeg.Price)
				fmt.Println("New Quantity of " + b + ":", replyVeg.Quantity)
			}

			return nil
		},
		Subcommands: []*cli.Command{
			{
			  Name:  "quantity",
			  Usage: "Update quantity of the vegitable",
			  Action: func(ctx *cli.Context) error {
					n := ctx.NArg()
					if n != 2 {
						return fmt.Errorf("Required arguments are not provided for update quantity operation")
					}

					b := ctx.Args().Get(0)
					c := ctx.Args().Get(1)

					quantity, _ := strconv.Atoi(c)

					a := common.UpdateVegetable{b, quantity}

					var replyVeg common.Vegetable
					var replyQuantity int
					client, err := rpc.DialHTTP("tcp", "localhost:9090")

					if err != nil {
						log.Fatal("Connection error: ", err)
					}

					erro := client.Call("MARKET.GetVegetableQuantity", b, &replyQuantity)
					if erro != nil {
						log.Fatal("Error: ", erro)
					}  else {
						fmt.Println("Old quantity of " + b + ":", replyQuantity)
					}

					erro2 := client.Call("MARKET.UpdateVegetableQuantity", a, &replyVeg)
					if erro2 != nil {
						log.Fatal("Error: ", erro2)
					}  else {
						fmt.Println("New Quantity of " + b + ":", replyVeg.Quantity)
						fmt.Println("Upated " + b + ": ", replyVeg)
					}

					return nil
			  },
			},
			{
			  Name:  "price",
			  Usage: "Update price of the vegitable",
			  Action: func(ctx *cli.Context) error {
				n := ctx.NArg()
				if n != 2 {
					return fmt.Errorf("Required arguments are not provided for update price operation")
				}

				b := ctx.Args().Get(0)
				c := ctx.Args().Get(1)

				price, _ := strconv.Atoi(c)

				a := common.UpdateVegetable{b, price}

				var replyVeg common.Vegetable
				var replyPrice int
				client, err := rpc.DialHTTP("tcp", "localhost:9090")

				if err != nil {
					log.Fatal("Connection error: ", err)
				}

				erro := client.Call("MARKET.GetVegetablePrice", b, &replyPrice)
				if erro != nil {
					log.Fatal("Error: ", erro)
				}  else {
					fmt.Println("Old price of " + b + ":", replyPrice)
				}
                
				erro2 := client.Call("MARKET.UpdateVegetablePrice", a, &replyVeg)
				if erro2 != nil {
					log.Fatal("Error: ", erro2)
				}  else {
					fmt.Println("New price of " + b + ":", replyVeg.Price)
					fmt.Println("Upated " + b + ": ", replyVeg)
				}
			
				return nil
		  },
			},
		  },
	}
}

func getAll(client *rpc.Client) {
	var replyList []string
	client.Call("MARKET.GetAllVegetables", "", &replyList)
	fmt.Println("Vegetables in Market: ", replyList)
}

func mainAction(ctx *cli.Context) error {
	ctx.App.Command("help").Run(ctx)

	return nil
}