package main

import (
	"fmt"
	"reflect"
	"syscall"
	"unsafe"
	"os"
	"net/http"
	"time"
	"bufio"
	"io"
	"strings"
	"strconv"
	"io/ioutil"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func main() {
	leSitesDoArquivo()
	for {
		exibeMenu()
	}
}

func imprimeLogs(){
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}

func iniciarMonitoramento(delay int){
	sites := []string{}
	sites = leSitesDoArquivo()

	fmt.Println("Iniciando Monitoramento")
	
	for i:=0; i<len(sites); i++{
		testaSite(sites[i])
		if i+1 != len(sites){
			time.Sleep(time.Duration(delay)*time.Second)
		}
	}
}

func leSitesDoArquivo() []string{
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil{
		fmt.Println("Ocorreu um erro!")
	}
	leitor := bufio.NewReader(arquivo)
	for{
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)// Remove spaces and \n
		
		sites = append(sites, linha)

		if err == io.EOF{
			break
		}
	}
	return sites
}

func testaSite(url string){
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Ocorreu um erro!")
	}
	switch response.StatusCode{
		case 200:
			fmt.Printf("Site: %s - carregado com sucesso !\n", url)
			registraLog(url, true)
		default:
			fmt.Printf("Site: %s - Erro %d\n", url, response.StatusCode)
			registraLog(url, false)
	}
}

func registraLog(site string, status bool){
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil{
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05")+" "+site + " - online: " + strconv.FormatBool(status)+"\n")
	arquivo.Close()
}

func leComando() int{
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func exibeInfo(){
	var name string = "Douglas"
	var version float64 = 1.1
	var age int64
	var undefiniedType = 9223372036854775807
	variavelCurta := 1.1
	fmt.Println("Hello, ", name, ", your age is: ", age, "\nVersion: ", version)
	fmt.Println("Var type: ", reflect.TypeOf(undefiniedType), ", var: ", undefiniedType)
	fmt.Println("Variável Curta tipo ", reflect.TypeOf(variavelCurta), ", Valor:", variavelCurta)
	terminalWidth := int(getWidth())
	for i := 0; i < terminalWidth; i++ {
		fmt.Printf("==")
	}
}

func exibeMenu(){
	fmt.Printf("%s\n%s\n%s\n",
		"1. Iniciar Monitoramento.",
		"2. Exibir Logs.",
		"0. Sair do Programa.")
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento(5)
			
		case 2:
			fmt.Println("Exibindo logs")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida")
			os.Exit(-1)
		}
}

func getWidth() uint {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return uint(ws.Col)
}

func devolveNomeEIdade() (string, int){
	nome := "Nome"
	idade := 100
	return nome, idade
}

/*
if comando == 1{
	fmt.Println("Iniciando Monitoramento")
} else if comando == 2{
	fmt.Println("Exibindo logs")
} else if comando == 0{
	fmt.Println("Saindo do programa...")
} else {
	fmt.Println("Opção inválida")
}*/