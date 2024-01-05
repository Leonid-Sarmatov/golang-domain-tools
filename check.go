package main

import (
	"fmt"
	//"log"
	"net"
	"time"
	//"net/netip"
)

func Check(destinatior string, port string) (string, bool) {
	// Создаем адрес запроса
	address := destinatior + ":" + port
	// Создаем максимальное время ожидния
	timeout := time.Duration(5 * time.Second)
	// Пытаемся подключиться по домену
	conn, err := net.DialTimeout("tcp", address, timeout)
	// Фалаг домена, поднят если и домен поднят
	flag := false

	var status string
	if err != nil {
		// Если ошибка, то пишем что домен не работает
		status = fmt.Sprintf("[DOWN] '%v' is unreachable,\nError: %v\n", destinatior, err)
	} else {
		defer conn.Close()
		// Иначе что домен работает
		status = fmt.Sprintf("[UP] '%v' is reachable,\nFrom: %+v\nTo: %+v\n",
			destinatior, addrToStringIP(conn.LocalAddr()), addrToStringIP(conn.RemoteAddr()))
		// Поднимаем флаг
		flag = true
	}
	return status, flag
}

func addrToStringIP(addr net.Addr) string {
	switch val := addr.(type) {
	case *net.TCPAddr:
		return fmt.Sprintf("%s:%d", val.IP, val.Port)
	default:
		return "Unknown address type"
	}
}

func GetListOfIPV4(destinatior string) ([]string, error) {
	// Создаем список с результатами
	listOfIP4 := make([]string, 0)
	// Получаем список всех адресов
	addresses, err := net.LookupHost(destinatior)
	if err != nil {
		return listOfIP4, err
	}

	// Оставляем только IPv4 и не нулевые адреса
	for _, val := range addresses {
		ip := net.ParseIP(val)
		if ip != nil && ip.To4() != nil {
			listOfIP4 = append(listOfIP4, val)
		}
	}
	return listOfIP4, nil
}
