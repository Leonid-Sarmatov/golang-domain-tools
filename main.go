package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	// Создаем экземпляр приложения
	app := &cli.App{
		// Задаем ему имя
		Name: "Healthchecker",
		// Задаем назначение
		Usage: "A tiny tool for checkgos whether a website is running or is down",
		// Задаем ключи для запуска в терминале
		Flags: []cli.Flag{
			// Первый флаг -d это домен
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to chek.",
				Required: true,
			},
			// Второй, необязательный, -p это порт прослушивания
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check.",
				Required: false,
			},
			// Третий, необязательный, -ip4 получение списка IPv4 адресов домена
			&cli.BoolFlag{
				Name:     "ip4-list",
				Aliases:  []string{"ip4"},
				Usage:    "Get a list of IP addresses. (only IPv4)",
				Required: false,
			},
		},
		// Задаем функцию выполнения
		Action: func(c *cli.Context) error {
			// Если ключа порта нет, то выставляем порт поумолчанию
			port := c.String("port")
			if port == "" {
				port = "80"
			}
			// Запускаем функцию проверки состояния домена
			status, isRun := Check(c.String("domain"), port)
			// Выводим статус
			fmt.Println(status)

			if isRun && c.IsSet("ip4-list") {
				list, err := GetListOfIPV4(c.String("domain"))
				if err != nil {
					fmt.Println("ERROR: ", err)
				} else {
					fmt.Println("IPv4 addresses:")
					for i, val := range list {
						fmt.Printf("%v. %v\n", i, val)
					}
				}
			}
			// Завершаем работу программы
			return nil
		},
	}

	// Спавн приложения
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
