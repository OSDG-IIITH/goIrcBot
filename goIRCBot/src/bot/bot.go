package main

import ("net"
        "log"
        "bufio"
        "fmt"
        "net/textproto"
        "strings"
        "functions"
      )

type Bot struct{
        server string
        port string
        nick string
        user string
        channel string
        pass string
        pread, pwrite chan string
        conn net.Conn
}

func NewBot() *Bot {
        return &Bot{server: "irc.freenode.net",
                    port: "6667",
                    nick: "teeheeBot",
                    channel: "#osdg-iiith",
                    pass: "",
                    conn: nil,
                    user: "blaze"}
}

func (bot *Bot) Connect() (conn net.Conn, err error){
  conn, err = net.Dial("tcp",bot.server + ":" + bot.port)
  if err != nil{
    log.Fatal("unable to connect to IRC server ", err)
  }
  bot.conn = conn
  log.Printf("Connected to IRC server %s (%s)\n", bot.server, bot.conn.RemoteAddr())
  return bot.conn, nil
}

func (bot *Bot) WriteMessage(message string){
  fmt.Fprintf(bot.conn, "PRIVMSG %s :%s\r\n", bot.channel, message)
}

func (bot *Bot) WriteMultMessage(vals ...string){
	for _, value:=range vals{
		bot.WriteMessage(value)
	}
}

func (bot *Bot) Pong(server string){
  fmt.Fprintf(bot.conn, "PONG %s", server)
  fmt.Printf("PONG %s", server)
}

func (bot *Bot) EvaluateLine(line string){

  splitUp := strings.Split(line, ":")
  
  if len(splitUp) > 2 {
    name := strings.Split(splitUp[1], "!")
    if strings.HasPrefix(splitUp[2], "!teehee ") {

      flags := strings.Split(splitUp[2], " ")

      if flags[1] == "help" {
      	bot.WriteMultMessage(functions.Help(name[0]))
      } else if flags[1] == "about" {
        bot.WriteMultMessage(functions.About(name[0]))
      } else if flags[1] == "rules" {
        bot.WriteMultMessage(functions.Rules(name[0]))
      } else if flags[1] == "commands" {
      	bot.WriteMultMessage(functions.Commands(name[0]))
      } else if flags[1] == "pastebin" {
      	bot.WriteMultMessage(functions.Pastebin(name[0]))
      } else if flags[1] == "stories" {
      	bot.WriteMultMessage(functions.Stories(name[0]))
      } else {
      	bot.WriteMessage(name[0]+": I didn't get that, try '!teehee help'")
      }
    }
  }
  if strings.HasPrefix(splitUp[0], "PING") {
    bot.Pong(splitUp[1])
  }
  fmt.Printf("%q\n", splitUp)
}

func main(){
  ircbot := NewBot()
  conn, _ := ircbot.Connect()
  
  fmt.Fprintf(conn, "USER %s 8 * :%s\r\n", ircbot.nick, ircbot.nick)
  fmt.Fprintf(conn, "NICK %s\r\n", ircbot.nick)
  fmt.Fprintf(conn, "JOIN %s\r\n", ircbot.channel) 
  defer conn.Close()
  
  reader := bufio.NewReader(conn)
  tp := textproto.NewReader( reader )
  for {
        line, err := tp.ReadLine()
        if err != nil {
            break // break loop on errors    
        }
        ircbot.EvaluateLine(line)
    }
}