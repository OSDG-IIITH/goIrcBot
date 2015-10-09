package functions

func About(name string) (string, string, string, string, string, string) {
	val1:="Hello @"+name+", Welcome to the Open Source Developers Group @ IIIT - H"
	val2:="Mailing List : https://groups.google.com/forum/?fromgroups#!forum/iiit-osdg"
	val3:="Blog : http://iiitosdg.wordpress.com/"
	val4:="IRC : Well, you guys are already here aren't you :P"
	val5:="GitHub : https://github.com/OSDG-IIITH/"
	val6:="IIIT's results at GSoC last summer: http://osdg.iiit.ac.in/gsoc15/"
	return val1, val2, val3, val4, val5, val6
}

func Commands(name string) (string, string, string, string, string, string, string, string) {
	val1:="@"+name+": here are some basic IRC commands:"
	val2:="'/join #channelname' (join channel of your choice | ex: /join #osdg-iiith, /join #ubuntu)\t'/chat nickname' (Opens a separate chatbox with another user, subject to confirmation. Chatbox is open even after server disconnects)"
	val3:="'/ignore nickname 3' (Use this to stop receiving messages from a particular user. Use '/ignore -r nickname 3' to unignore that user)\t'/help' (Opens up the built-in IRC help menu)"
	val4:=" '(message) /me (message)' (action message | ex: If your nick is 'john', then /me 'likes marshmallows' will be displayed as 'john likes marshmallows'.)\t'/msg nickname message' (PMs another user on the channel, opens up a private window. It is good etiquette to ask the concerned person if you can pm him or her first.)"
	val5:="'/nick newnickname' (Changes your nickname to the new one you specify | '/nick gandalf' will change your nick to 'gandalf')\t'/notice nickname message' (A notice is used to send a short message to another person without opening up a private window.)"
	val6:="'/part' and '/partall' (Helps you leave one and all channels respectively.)\t'/ping nicname' (Gives you the ping time or lag time between you and the person you pinged.)"
	val7:="'/query nickname' (Similar to the /msg command, but forces chat window to pop open. Poor etiquette to use this for strangers.)\t'/quit <message>' (Command to exit IRC altogether)"
	val8:="'/whois nickname' (This command helps you know more about a particular user.)"
	return val1, val2, val3, val4, val5, val6, val7, val8
}

func Help(name string) (string, string, string) {
	val1:="Hi @"+name+"!, I am teeheebot, and I've been written in golang. My functions are:"
	val2:="about    commands    help    pastebin    rules    stories"
	val3:="For usage of the above: !teehee <function>" 
	return val1, val2, val3
}

func Rules(name string) (string, string, string, string, string) {	
	val1:="Hello "+name+"! Here are some basic informal rules you should keep in mind:" 
	val2:="Rule One: Google. Not kidding. To know what this truly means, visit: https://iiitosdg.wordpress.com/2013/08/10/searching-for-an-answer/"
	val3:="Rule Two: Don't ask to ask, just ask. Read up at http://sol.gfxile.net/dontask.html"
	val4:="Rule Three: Respect netiquette. Find out more at http://www.ircbeginner.com/ircinfo/etiquette.html"
	val5:="The rest of them, you'll figure as you go along. :P Refer to '!teehee help' for more info."
	return val1, val2, val3, val4, val5
}

func Pastebin(name string) (string, string) {
	val1:="Hello @"+name+", here are some pastebin links you can use:"
	val2:="http://paste.ubuntu.com/\thttp://pastebin.com\thttp://ipaste.eu\thttp://fpaste.org/"
	return val1, val2	
}

func Stories(name string) (string, string, string) {
	val1:="OSDG-IIITH consists of amazing people, each with an inspiring story about their individual open source journey. :)"
	val2:="Right now, we have accounts by Amogh, Anurag, Arushi, Kalpit, Mrinal and Vatika."
	val3:="Go ahead! Read them at http://osdg.iiit.ac.in/stories.php"
	return val1, val2, val3
}
