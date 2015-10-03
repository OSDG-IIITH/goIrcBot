package main 

import ("fmt"
	"net"
	"log"
	"bufio"
	"net/textproto" //Is it not included in the "net" package?
)

type Bot struct {
	server string  
	port string
	nick string
	user string //??
	channel string
	pass string //??
	pread, pwrite chan string //What data are you passing around between which goroutines?
	conn net.Conn //Is net.Conn a datatype? What kind of value will conn be storing?
}

func NewBot() *Bot { //Why not return a value? Why the reference?
	return &Bot{
		server: "irc.freenode.net",
		port: "6667",
		nick: "teeheeBot",
		channel: "#osdg-iiith",
		pass: "",
		conn: nil,
		user: "blaze" //Who/What is blaze? The objective here?
	}
}

func (bot *Bot) Connect() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp",bot.server+":"+bot.port) //Why TCP? 
	if err!=nil {
		log.Fatal("Cannot connect to server now. :/", err) //Will this message be passed to an 'stdout' here before sys.exit(1)?
	}
	bot.conn=conn
	log.Printf("Connected to IRC server %s (%s)\n", bot.server, bot.conn.RemoteAddr()) //Ummm, what is a Remote Address? *goes and hides away embarrassment*
	return bot.conn, nil
}

func (bot *Bot) Pong(server string){ //What is a Pong message?
	fmt.Fprintf(bot.conn,"PONG %s",server)//not bot.server?
	fmt.Printf("PONG %s", server)
}

func (bot *Bot) WriteMessage(messsage string) {
	fmt.Fprintf(bot.conn, "PRIVMSG %s :%s\r\n", bot.channel, message)
}

func (bot *Bot) EvaluateLine(line string){
	splitUp:=strings.Split(line,":")//What kind of data are we looking at in line here? Why is the colon a delimiter for the future list?

	if len(splitUp)>2 { //What are the list elements if true? 
		name:=strings.Split(splitUp[1],"!")
		if strings.HasPrefix(splitUp[2], "!teehee ") {
			flags:=strings.Split(splitUp[2]," ")
			if flags[1]=="help"{
				bot.WriteMessage(name[0]+": TEEHEEBOT : written in Golang.")//name[0]??
        		bot.WriteMessage(name[0]+": Current Functions: help, about")
        		bot.WriteMessage(name[0]+": USAGE: '!teehee <function> [flags]'") 
			} else if flags[1]=="about" {
				bot.WriteMessage(name[0]+": Open Source Developers Group @ IIIT - H")
		        bot.WriteMessage(name[0]+": Mailing List : https://groups.google.com/forum/?fromgroups#!forum/iiit-osdg")
				bot.WriteMessage(name[0]+": Blog : http://iiitosdg.wordpress.com/")
		        bot.WriteMessage(name[0]+": IRC : Well, you guys are already here aren't you :P")
		        bot.WriteMessage(name[0]+": GitHub : https://github.com/OSDG-IIITH/")
		        bot.WriteMessage(name[0]+": Want to get a project forked under the github group? Register it at http://osdg.iiit.ac.in/github/")
		        bot.WriteMessage(name[0]+": Doing GSoC this summer? Check out http://osdg.iiit.ac.in/gsoc15/") 
			} else  {
				bot.WriteMessage(name[0]+": I didn't get that, try '!teehee help' ??")
			}
		}
	}
	if strings.HasPrefix(splitUp[0],"PING") {//??
		bot.Pong(splitUp[1])
	}
	fmt.Printf("%q\n", splitUp)
}


func main() {
	ircbot:=newbot()
	conn,_:=ircbot.Connect()

	fmt.Fprintf(conn,"USER %s 8 * :%s\r\n", ircbot.nick, ircbot.nick)//??
	fmt.Fprintf(conn,"NICK %s\r\n", ircbot.nick)//??
	fmt.Fprintf(conn,"JOIN %s\r\n", ircbot.channel)//??
	defer conn.Close()//Since the for loop below never ends, I assume this is here just for formality and structure?

	reader:=bufio.NewReader(conn)//Creates a buffer with size of net.Conn?
	tp:=textproto.NewReader(reader)//https://golang.org/pkg/net/textproto/#NewReader :This is to stop DOS attacks? To bound buffer size?
	for { //Infinite Loop? What kind of errors can tp.ReadLine() generate? When?
		line,err:=tp.ReadLine()
		if err!=nil{
			break
		}
		ircbot.EvaluateLine()
	}
}
//I think I need to study IRC client protocol first, understand how IRC networks, what sockets and servers are, what the different commands are.
//Question that exposes my non-existent knowledge of networks: Why doesn't the bot connect to the server when run from the terminal? :P
//Note of Appreciation: Amogh Pradeep, you are a genius. :P
