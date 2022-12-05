package teleport

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"
	datamodels "teleport-client/datamodel"
	"teleport-client/utility"
)

func TestTeleportSSH(config utility.Teleport) datamodels.Stats {

	cmd := exec.Command(config.Tsh, config.TshCommand, config.TeleportProxy, config.Identity, config.InferenceServer, config.SSHCommand)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("!!!!!!! Error while tring to init a output pipe ")
		log.Print("!!! Error >>")
		panic(err)
	}
	stdErrPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Println("!!!!!!! Error while tring to init a error pipe ")
		log.Print("!!! Error >>")
		panic(err)
	}
	err = cmd.Start()
	if err != nil {
		log.Println("!!!!!!! Error while tyring to start the cmd")
		log.Print("!!! Error >>")
		panic(err)
	}
	cmdOp, err := io.ReadAll(stdoutPipe)
	if err != nil {
		log.Println("!!!!!!! Error while tring to process command output", "err>>", err)
		log.Print("!!! Error >>")
		panic(err)
	}
	cmdErr, err := io.ReadAll(stdErrPipe)
	if err != nil {
		log.Println("!!!!!!! Error while tring to process command Error", "err>>", err)
		log.Print("!!! Error >>")
		panic(err)
	}
	log.Println(string(cmdErr))
	cmd.Wait()
	err = writeOpToTxtFile(string(cmdOp), "status/edgeFolderUpdated.json")
	if err != nil {
		log.Println("Error while trying to write date to the file")
		panic(err)
	}
	return readfile("status/edgeFolderUpdated.json")
}

func writeOpToTxtFile(op, name string) error {
	log.Println("Creating file with name: ", name)
	file, err := os.Create(name)
	if err != nil {
		log.Println("Error while trying to create file with name: ")
		log.Print("!!! Error >>")
		return err
	}
	opb := []byte(op)
	noOfBytesWritten, err := file.Write(opb)
	if noOfBytesWritten < len(opb) || err != nil {
		log.Println("Program was not able to write complete o/p to the file")
		log.Print("!!! Error >>")
		return err
	}
	return file.Close()
}

func readfile(path string) datamodels.Stats {

	data, err := os.ReadFile(path)
	if err != nil {
		log.Println("Error while trying to read the data from file: ", path)
		log.Println("Error: ", err)
	}
	stats := datamodels.Stats{}

	err = json.Unmarshal(data, &stats)
	if err != nil {
		log.Println("Error while trying unmarshal the json data")
		log.Println("Error: ", err)
	}
	return stats

	// edgeDuration := make(map[string]float64)
	// file, err := os.Open(path)
	// if err != nil {
	// 	fmt.Println("err while opening file:", err)
	// 	return nil
	// }
	// defer file.Close()

	// sc := bufio.NewScanner(file)

	// if err := sc.Err(); err != nil {
	// 	log.Fatal("Error while scanning the file: ", err)
	// 	return nil
	// }

	// location, err := time.LoadLocation("CET")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// tNow := time.Now().In(location)
	// if tNow.Hour() > 7 && tNow.Hour() < 22 {
	// 	for sc.Scan() {
	// 		edgeFolderData := sc.Text()
	// 		timeName := strings.Split(edgeFolderData, "\t")
	// 		t := strings.Split(timeName[0], ":")
	// 		name := strings.ReplaceAll(timeName[1], " ", "")

	// 		// fmt.Println(name, "->", t)
	// 		h, err := strconv.Atoi(t[0])
	// 		if err != nil {
	// 			fmt.Println("err while converting hour from string to int:", err)
	// 		}
	// 		m, err := strconv.Atoi(strings.ReplaceAll(t[1], " ", ""))
	// 		if err != nil {
	// 			fmt.Println("err while converting min from string to int:", err)
	// 		}

	// 		tEdge := time.Date(tNow.Year(), tNow.Month(), tNow.Day(), h, m, 0, 0, location)
	// 		tSub := tNow.Sub(tEdge).Minutes()

	// 		edgeDuration[name] = tSub

	// 		// fmt.Println("-------Now---------->", tNow)
	// 		// fmt.Println("-------Edge---------->", tEdge)

	// 		// fmt.Println("-----------Dif Now-Edge------->", tNow.Sub(tEdge).Minutes())
	// 	}
	// 	return edgeDuration
	// } else {
	// 	return nil
	// }
}
