## 1.时代的里程碑——即时通信

前阵子看了《创业时代》,电视剧的剧情大概是这样的:IT工程师郭鑫年与好友罗维与投行精英那蓝等人一起，踏上互联网创业之路。创业开发的是一款叫做“魔晶”的IM产品。郭鑫年在第一次创业失败后，离了婚，还欠了很多外债，骑着自行车经历了西藏一次生死诀别之后产生了灵感，想要创作一款IM产品“魔晶”，“魔晶”的初衷是为了增加人与人之间的感情，虽然剧情纯属虚构，但确实让人浮想QQ当初的设想是不是就是这样的呢？

有一点是可以确定的，即时通信确实是一个时代的里程碑。腾讯的强大离不开两款产品:QQ和微信,这两款产品设计的思路是不一样的，QQ依托于IM系统，为了打造个人空间、全民娱乐而设计，我们常常会看到QQ被初高中生喜爱，QQ账号也往往与音乐、游戏绑定在一起；微信从QQ导流以后，主打商业领域，从刚开始推出微信支付与支付宝竞争，在商业支付领域占得了一席之地(微信支付主要被用户用于小额支付场景，支付宝主要用在企业大额转账、个人金融理财领域)以后。微信又相继推出了公众号、小程序，很明显在商业领域已经占据了支付宝的上风，成为了商业APP中的霸主，后来才有了聊天宝、多闪和马桶三大门派围攻微信的闹剧，结果大家可能都知道了......

![](https://user-gold-cdn.xitu.io/2020/1/12/16f9a3aef634edfc?w=720&h=405&f=png&s=516840)

即使这样，也不可低估了支付宝在商业领域的价值。和微信产品设计的初衷不同，支付宝更青睐于拓展功能自己来集成与实现，做的比较精致。支付宝在支付安全性方面，做的比微信好很多，整个应用用起来也比微信要顺畅，支付宝也有自己的小程序，不过往往都是和相关企业合作或打通接口来创建应用，比如生活缴费、饿了么外卖、滴滴出行等等。而微信则更多将应用的创建权限开放给开发者，由企业开发者来创建小程序、维护公众号，从而实现自己的商业价值，事实证明，微信很成功!

阿里依托于IM系统进击办公领域，打造了“钉钉”,又是一款比较精致的产品，其中打卡考勤、请假审批、会议管理都做的非常好，和微信不同的是，企业通过钉钉交流的信息，对方是能看到信息是否“已读”的(毕竟是办公，这个功能还是很有必要的)。腾讯也不甘示弱，创建“企业微信”，开始和“钉钉”正面交锋，虽然在市场份额上还是落后于钉钉，但用户增长很快。


![](https://user-gold-cdn.xitu.io/2020/1/12/16f9a689e8d15c4e?w=793&h=297&f=png&s=124473)

>钉钉于2015年1月正式上线，2016年4月腾讯正式发布企业微信1.0版本，也只有简单的考勤、请假、报销等功能，在产品功能上略显平淡。彼时再看钉钉，凭借先发优势，初期就确定的产品线“讨好”老板，2016年企业数100万，2018年这个数量上升到700万，可见钉钉发展速度之快，稳固了钉钉在B端市场的地位。企业微信早期举棋不定的打法，也让它在企业OA办公上玩不过钉钉。但企业微信在发布3.0版本后，局面开始扭转，钉钉在用户数量上似乎已经饱和，难以有新的突破，而企业微信才真正开始逐渐占据市场。

依托于IM系统发展起来的企业还有陌陌、探探，相比较与微信来讲，它们的功能更集中于交友和情感。(不知道这是不是人家企业每年年终都人手一部iphone的原因，开个玩笑)

笔者今年参加了一次Gopher大会，有幸听探探的架构师分享了它们今年微服务化的过程，本文快速搭建的IM系统也是使用Go语言来快速实现的，这里先和各位分享一下探探APP的架构图:

![](https://user-gold-cdn.xitu.io/2020/1/12/16f9a63991fc4485?w=1924&h=1160&f=png&s=529064)

以上讲了一些IM系统的产品的设计，下边我们回归主题，大概说一下本文的章节内容安排。

## 2.章节概述

本文的目的是帮助读者较为深入的理解socket协议,并快速搭建一个高可用、可拓展的IM系统(文章标题纯属引人眼球，不是真的，请读者不要在意。)，同时帮助读者了解IM系统后续可以做哪些优化和改进。麻雀虽小，五脏俱全，该IM系统包含基本的注册、登录、添加好友基础功能，另外提供单聊、群聊，并且支持发送文字、表情和图片，在搭建的系统上，读者可轻松的拓展语音、视频聊天、发红包等业务。为了帮助读者更清楚的理解IM系统的原理,第3节我会专门深入讲解一下websocket协议，websocket是长链接中比较常用的协议；然后第4节会讲解快速搭建IM系统的技巧和主要代码实现；在第5节笔者会对IM系统的架构升级和优化提出一些建议和思路；在最后章节做本文的回顾总结。

## 3.深入理解websocket协议

Web Sockets的目标是在一个单独的持久连接上提供全双工、双向通信。在Javascript创建了Web Socket之后，会有一个HTTP请求发送到浏览器以发起连接。在取得服务器响应后，建立的连接会将HTTP升级从HTTP协议交换为WebSocket协议。
由于WebSocket使用自定义的协议，所以URL模式也略有不同。未加密的连接不再是http://，而是ws://;加密的连接也不是https://，而是wss://。在使用WebSocket URL时，必须带着这个模式，因为将来还有可能支持其他的模式。
使用自定义协议而非HTTP协议的好处是，能够在客户端和服务器之间发送非常少量的数据，而不必担心HTTP那样字节级的开销。由于传递的数据包很小，所以WebSocket非常适合移动应用。
上文中只是对Web Sockets进行了笼统的描述，接下来的篇幅会对Web Sockets的细节实现进行深入的探索，本文接下来的四个小节不会涉及到大量的代码片段，但是会对相关的API和技术原理进行分析，相信大家读完下文之后再来看这段描述，会有一种豁然开朗的感觉。

### 3.1 WebSocket复用了HTTP的握手通道

“握手通道”是HTTP协议中客户端和服务端通过"TCP三次握手"建立的通信通道。客户端和服务端使用HTTP协议进行的每次交互都需要先建立这样一条“通道”，然后通过这条通道进行通信。我们熟悉的ajax交互就是在这样一个通道上完成数据传输的，只不过ajax交互是短连接，在一次request->response之后，“通道”连接就断开了。下面是HTTP协议中建立“握手通道”的过程示意图：

![](https://user-gold-cdn.xitu.io/2020/1/13/16f9ef56adac8cc9?w=800&h=478&f=png&s=163832)

上文中我们提到：在Javascript创建了WebSocket之后，会有一个HTTP请求发送到浏览器以发起连接，然后服务端响应，这就是“握手“的过程，在这个握手的过程当中，客户端和服务端主要做了两件事情：

- 建立了一条连接“握手通道”用于通信（这点和HTTP协议相同，不同的是HTTP协议完成数据交互后就释放了这条握手通道，这就是所谓的“短连接”，它的生命周期是一次数据交互的时间，通常是毫秒级别的。）
- 将HTTP协议升级到WebSocket协议，并复用HTTP协议的握手通道，从而建立一条持久连接。
说到这里可能有人会问：HTTP协议为什么不复用自己的“握手通道”，而非要在每次进行数据交互的时候都通过TCP三次握手重新建立“握手通道”呢？答案是这样的：虽然“长连接”在客户端和服务端交互的过程中省去了每次都建立“握手通道”的麻烦步骤，但是维持这样一条“长连接”是需要消耗服务器资源的，而在大多数情况下，这种资源的消耗又是不必要的，可以说HTTP标准的制定经过了深思熟虑的考量。到我们后边说到WebSocket协议数据帧时，大家可能就会明白，维持一条“长连接”服务端和客户端需要做的事情太多了。

说完了握手通道，我们再来看HTTP协议如何升级到WebSocket协议的。

### 3.2 HTTP协议升级为WebSocket协议

升级协议需要客户端和服务端交流，服务端怎么知道要将HTTP协议升级到WebSocket协议呢？它一定是接收到了客户端发送过来的某种信号。下面是我从谷歌浏览器中截取的“客户端发起协议升级请求的报文”，通过分析这段报文，我们能够得到有关WebSocket中协议升级的更多细节。


![](https://user-gold-cdn.xitu.io/2020/1/13/16f9ef77ca1e6476?w=800&h=387&f=png&s=125303)

首先，客户端发起协议升级请求。采用的是标准的HTTP报文格式，且只支持GET方法。下面是重点请求的首部的意义：

- Connection：Upgrade：表示要升级的协议
- Upgrade: websocket：表示要升级到websocket协议
- Sec-WebSocket-Version: 13：表示websocket的版本
- Sec-WebSocket-Key:UdTUf90CC561cQXn4n5XRg== ：与Response Header中的响应首部Sec-WebSocket-Accept: GZk41FJZSYY0CmsrZPGpUGRQzkY=是配套的，提供基本的防护，比如恶意的连接或者无意的连接。其中Connection就是我们前边提到的，客户端发送给服务端的信号，服务端接受到信号之后，才会对HTTP协议进行升级。那么服务端怎样确认客户端发送过来的请求是否是合法的呢？在客户端每次发起协议升级请求的时候都会产生一个唯一码：Sec-WebSocket-Key。服务端拿到这个码后，通过一个算法进行校验，然后通过Sec-WebSocket-Accept响应给客户端，客户端再对Sec-WebSocket-Accept进行校验来完成验证。这个算法很简单：

1. 将Sec-WebSocket-Key跟全局唯一的（GUID，[RFC4122]）标识：258EAFA5-E914-47DA-95CA-C5AB0DC85B11拼接
2. 通过SHA1计算出摘要，并转成base64字符串

258EAFA5-E914-47DA-95CA-C5AB0DC85B11这个字符串又叫“魔串"，至于为什么要使用它作为Websocket握手计算中使用的字符串，这点我们无需关心，只需要知道它是RFC标准规定就可以了，官方的解析也只是简单的说此值不大可能被不明白WebSocket协议的网络终端使用。我们还是用世界上最好的语言来描述一下这个算法吧。

```
public function dohandshake($sock, $data, $key) {
        if (preg_match("/Sec-WebSocket-Key: (.*)\r\n/", $data, $match)) {
            $response = base64_encode(sha1($match[1] . '258EAFA5-E914-47DA-95CA-C5AB0DC85B11', true));
            $upgrade  = "HTTP/1.1 101 Switching Protocol\r\n" .
                "Upgrade: websocket\r\n" .
                "Connection: Upgrade\r\n" .
                "Sec-WebSocket-Accept: " . $response . "\r\n\r\n";
            socket_write($sock, $upgrade, strlen($upgrade));
            $this->isHand[$key] = true;
        }
    }
```

服务端响应客户端的头部信息和HTTP协议的格式是相同的，HTTP1.1协议是以换行符(\r\n)分割的，我们可以通过正则匹配解析出Sec-WebSocket-Accept的值，这和我们使用curl工具模拟get请求是一个道理。这样展示结果似乎不太直观，我们使用命令行CLI来根据上图中的Sec-WebSocket-Key和握手算法来计算一下服务端返回的Sec-WebSocket-Accept是否正确：

![](https://user-gold-cdn.xitu.io/2020/1/13/16f9ef9e287782cc?w=800&h=102&f=png&s=46725)

从图中可以看到，通过算法算出来的base64字符串和Sec-WebSocket-Accept是一样的。那么假如服务端在握手的过程中返回一个错误的Sec-WebSocket-Accept字符串会怎么样呢？当然是客户端会报错，连接会建立失败，大家可以尝试一下，例如将全局唯一标识符258EAFA5-E914-47DA-95CA-C5AB0DC85B11改为258EAFA5-E914-47DA-95CA-C5AB0DC85B12。

### 3.3 WebSocket的帧和数据分片传输

下图是我做的一个测试：将小说《飘》的第一章内容复制成文本数据，通过客户端发送到服务端，然后服务端响应相同的信息完成了一次通信。

![](https://user-gold-cdn.xitu.io/2020/1/13/16f9efacb4a7b73a?w=800&h=292&f=png&s=159300)

可以看到一篇足足有将近15000字节的数据在客户端和服务端完成通信只用了150ms的时间。我们还可以看到浏览器控制台中frame栏中显示的客户端发送和服务端响应的文本数据，你一定惊讶WebSocket通信强大的数据传输能力。数据是否真的像frame中展示的那样客户端直接将一大篇文本数据发送到服务端，服务端接收到数据之后，再将一大篇文本数据返回给客户端呢？这当然是不可能的，我们都知道HTTP协议是基于TCP实现的，HTTP发送数据也是分包转发的，就是将大数据根据报文形式分割成一小块一小块发送到服务端,服务端接收到客户端发送的报文后，再将小块的数据拼接组装。关于HTTP的分包策略，大家可以查看相关资料进行研究，websocket协议也是通过分片打包数据进行转发的，不过策略上和HTTP的分包不一样。frame（帧）是websocket发送数据的基本单位，下边是它的报文格式：


![](https://user-gold-cdn.xitu.io/2020/1/13/16f9efb2f2308bb5?w=800&h=588&f=png&s=509398)

报文内容中规定了数据标示,操作代码、掩码、数据、数据长度等格式。不太理解没关系，下面我通过讲解大家只要理解报文中重要标志的作用就可以了。首先我们明白了客户端和服务端进行Websocket消息传递是这样的:

- 客户端：将消息切割成多个帧，并发送给服务端。
- 服务端：接收消息帧，并将关联的帧重新组装成完整的消息。

服务端在接收到客户端发送的帧消息的时候，将这些帧进行组装，它怎么知道何时数据组装完成的呢？这就是报文中左上角FIN(占一个比特)存储的信息，1表示这是消息的最后一个分片（fragment)如果是0,表示不是消息的最后一个分片。websocket通信中，客户端发送数据分片是有序的，这一点和HTTP不一样，HTTP将消息分包之后，是并发无序的发送给服务端的，包信息在数据中的位置则在HTTP报文中存储，而websocket仅仅需要一个FIN比特位就能保证将数据完整的发送到服务端。
接下来的RSV1,RSV2,RSV3三个比特位的作用又是什么呢？这三个标志位是留给客户端开发者和服务端开发者开发过程中协商进行拓展的，默认是0。拓展如何使用必须在握手的阶段就协商好，其实握手本身也是客户端和服务端的协商。

### 3.4 Websocket连接保持和心跳检测

Websocket是长连接，为了保持客户端和服务端的实时双向通信，需要确保客户端和服务端之间的TCP通道保持连接没有断开。但是对于长时间没有数据往来的连接，如果依旧保持着，可能会浪费服务端资源。但是不排除有些场景，客户端和服务端虽然长时间没有数据往来，仍然需要保持连接，就比如说你几个月没有和一个QQ好友聊天了，突然有一天他发QQ消息告诉你他要结婚了，你还是能在第一时间收到。那是因为，客户端和服务端一直再采用心跳来检查连接。客户端和服务端的心跳连接检测就像打乒乓球一样：

- 发送方->接收方：ping
- 接收方->发送方：pong

等什么时候没有ping、pong了，那么连接一定是存在问题了。
说了这么多，接下来我使用Go语言来实现一个心跳检测，Websocket通信实现细节是一件繁琐的事情，直接使用开源的类库是比较不错的选择，我使用的是：gorilla/websocket。这个类库已经将websocket的实现细节（握手，数据解码)封装的很好啦。下面我就直接贴代码了:

```
package main

import (
    "net/http"
    "time"

    "github.com/gorilla/websocket"
)

var (
    //完成握手操作
    upgrade = websocket.Upgrader{
       //允许跨域(一般来讲,websocket都是独立部署的)
       CheckOrigin:func(r *http.Request) bool {
            return true
       },
    }
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
   var (
         conn *websocket.Conn
         err error
         data []byte
   )
   //服务端对客户端的http请求(升级为websocket协议)进行应答，应答之后，协议升级为websocket，http建立连接时的tcp三次握手将保持。
   if conn, err = upgrade.Upgrade(w, r, nil); err != nil {
        return
   }

    //启动一个协程，每隔1s向客户端发送一次心跳消息
    go func() {
        var (
            err error
        )
        for {
            if err = conn.WriteMessage(websocket.TextMessage, []byte("heartbeat")); err != nil {
                return
            }
            time.Sleep(1 * time.Second)
        }
    }()

   //得到websocket的长链接之后,就可以对客户端传递的数据进行操作了
   for {
         //通过websocket长链接读到的数据可以是text文本数据，也可以是二进制Binary
        if _, data, err = conn.ReadMessage(); err != nil {
            goto ERR
     }
     if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
         goto ERR
     }
   }
ERR:
    //出错之后，关闭socket连接
    conn.Close()
}

func main() {
    http.HandleFunc("/ws", wsHandler)
    http.ListenAndServe("0.0.0.0:7777", nil)
}
```

借助go语言很容易搭建协程的特点，我专门开启了一个协程每秒向客户端发送一条消息。打开客户端浏览器可以看到，frame中每秒的心跳数据一直在跳动，当长链接断开之后，心跳就没有了，就像人没有了心跳一样：

![](https://user-gold-cdn.xitu.io/2020/1/13/16f9efd3635a0cdb?w=800&h=610&f=png&s=166996)

大家对websocket协议已经有了了解，接下来就让我们一起快速搭建一个高性能、可拓展的IM系统吧。

## 4.快速搭建高性能、可拓展的IM系统

### 4.1 系统架构和代码文件目录结构

下图是一个比较完备的IM系统架构：包含了C端、接入层（通过协议接入）、S端处理逻辑和分发消息、存储层用来持久化数据。

![](https://user-gold-cdn.xitu.io/2020/1/18/16fb695286d33b06?w=471&h=238&f=png&s=50062)

我们本节C端使用的是Webapp, 通过Go语言渲染Vue模版快速实现功能；接入层使用的是websocket协议，前边已经进行了深入的介绍；S端是我们实现的重点，其中鉴权、登录、关系管理、单聊和群聊的功能都已经实现，读者可以在这部分功能的基础上再拓展其他的功能，比如：视频语音聊天、发红包、朋友圈等业务模块；存储层我们做的比较简单，只是使用Mysql简单持久化存储了用户关系，然后聊天中的图片资源我们存储到了本地文件中。虽然我们的IM系统实现的比较简化，但是读者可以在次基础上进行改进、完善、拓展，依然能够作出高可用的企业级产品。

我们的系统服务使用Go语言构建，代码结构比较简洁，但是性能比较优秀（这是Java和其他语言所无法比拟的），单机支持几万人的在线聊天。

下边是代码文件的目录结构:

```
app
│   ├── args
│   │   ├── contact.go
│   │   └── pagearg.go
│   ├── controller           //控制器层，api入口
│   │   ├── chat.go
│   │   ├── contract.go
│   │   ├── upload.go
│   │   └── user.go
│   ├── main.go             //程序入口
│   ├── model               //数据定义与存储
│   │   ├── community.go
│   │   ├── contract.go
│   │   ├── init.go
│   │   └── user.go
│   ├── service             //逻辑实现
│   │   ├── contract.go
│   │   └── user.go
│   ├── util                //帮助函数    
│   │   ├── md5.go
│   │   ├── parse.go
│   │   ├── resp.go
│   │   └── string.go
│   └── view                //模版资源
│   │   ├── ...
asset                       //js、css文件
resource                    //上传资源，上传图片会放到这里
```

从入口函数main.go开始，我们定义了controller层，是客户端api的入口。service用来处理主要的用户逻辑，消息分发、用户管理都在这里实现。model层定义了一些数据表，主要是用户注册和用户好友关系、群组等信息，存储到mysql。util包下是一些帮助函数，比如加密、请求响应等。view下边存储了模版资源信息，上边所说的这些都在app文件夹下存储，外层还有asset用来存储css、js文件和聊天中会用到的表情图片等。resource下存储用户聊天中的图片或者视频等文件。总体来讲，我们的代码目录机构还是比较简洁清晰的。

了解了我们要搭建的IM系统架构，我们再来看一下架构重点实现的功能吧。

### 4.2 10行代码万能模版渲染

Go语言提供了强大的html渲染能力，非常简单的构建web应用，下边是实现模版渲染的代码，它太简单了，以至于可以直接在main.go函数中实现:

```
func registerView() {
	tpl, err := template.ParseGlob("./app/view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplName := v.Name()
		http.HandleFunc(tplName, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplName, nil)
		})
	}
}
...
func main() {
    ......
    http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.Handle("/resource/", http.FileServer(http.Dir(".")))
	registerView()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
```

Go实现静态资源服务器也很简单，只需要调用http.FileServer就可以了，这样html文件就可以很轻松的访问依赖的js、css和图标文件了。使用http/template包下的ParseGlob、ExecuteTemplate又可以很轻松的解析web页面，这些工作完全不依赖与nginx。现在我们就完成了登录、注册、聊天C端界面的构建工作:

<center class="half">
    <img src="https://user-gold-cdn.xitu.io/2020/1/18/16fb8326841d2b07?w=568&h=956&f=png&s=32195" width="200"/>
    <img src="https://user-gold-cdn.xitu.io/2020/1/18/16fb8329af742ad2?w=566&h=954&f=png&s=54763" width="200"/><img src="图片链接" width="200"/>
</center>

### 4.3 注册、登录和鉴权

之前我们提到过，对于注册、登录和好友关系管理，我们需要有一张`user`表来存储用户信息。我们使用`github.com/go-xorm/xorm`来操作`mysql`,首先看一下`mysql`表的设计:

app/model/user.go
```
package model

import "time"

const (
	SexWomen = "W"
	SexMan = "M"
	SexUnknown = "U"
)

type User struct {
	Id         int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`
	Mobile   string 		`xorm:"varchar(20)" form:"mobile" json:"mobile"`
	Passwd       string	`xorm:"varchar(40)" form:"passwd" json:"-"`   // 用户密码 md5(passwd + salt)
	Avatar	   string 		`xorm:"varchar(150)" form:"avatar" json:"avatar"`
	Sex        string	`xorm:"varchar(2)" form:"sex" json:"sex"`
	Nickname    string	`xorm:"varchar(20)" form:"nickname" json:"nickname"`
	Salt       string	`xorm:"varchar(10)" form:"salt" json:"-"`
	Online     int	`xorm:"int(10)" form:"online" json:"online"`   //是否在线
	Token      string	`xorm:"varchar(40)" form:"token" json:"token"`   //用户鉴权
	Memo      string	`xorm:"varchar(140)" form:"memo" json:"memo"`
	Createat   time.Time	`xorm:"datetime" form:"createat" json:"createat"`   //创建时间, 统计用户增量时使用
}
```

我们`user`表中存储了用户名、密码、头像、用户性别、手机号等一些重要的信息，比较重要的是我们也存储了token标示用户在用户登录之后，http协议升级为websocket协议进行鉴权，这个细节点我们前边提到过，下边会有代码演示。接下来我们看一下model初始化要做的一些事情吧：

app/model/init.go
```
package model

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	dsnName := "root:root@(127.0.0.1:3306)/chat?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, dsnName)
	if err != nil && err.Error() != ""{
		log.Fatal(err)
	}
	DbEngine.ShowSQL(true)
	//设置数据库连接数
	DbEngine.SetMaxOpenConns(10)
	//自动创建数据库
	DbEngine.Sync(new(User), new(Community), new(Contact))

	fmt.Println("init database ok!")
}
```

我们创建一个DbEngine全局mysql连接对象，设置了一个大小为10的连接池。model包里的init函数在程序加载的时候会先执行，对Go语言熟悉的同学应该知道这一点。我们还设置了一些额外的参数用于调试程序，比如：设置打印运行中的sql，自动的同步数据表等，这些功能在生产环境中可以关闭。我们的model初始化工作就做完了，非常简陋，在实际的项目中，像数据库的用户名、密码、连接数和其他的配置信息，建议设置到配置文件中，然后读取，而不像本文硬编码的程序中。

注册是一个普通的api程序，对于Go语言来说，完成这件工作太简单了，我们来看一下代码:

```
############################
//app/controller/user.go
############################
......
//用户注册
func UserRegister(writer http.ResponseWriter, request *http.Request) {
	var user model.User
	util.Bind(request, &user)
	user, err := UserService.UserRegister(user.Mobile, user.Passwd, user.Nickname, user.Avatar, user.Sex)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
	}
}
......
############################
//app/service/user.go
############################
......
type UserService struct{}

//用户注册
func (s *UserService) UserRegister(mobile, plainPwd, nickname, avatar, sex string) (user model.User, err error) {
    registerUser := model.User{}
    _, err = model.DbEngine.Where("mobile=? ", mobile).Get(&registerUser)
    if err != nil {
    	return registerUser, err
	}
	//如果用户已经注册,返回错误信息
	if registerUser.Id > 0 {
		return registerUser, errors.New("该手机号已注册")
	}

	registerUser.Mobile = mobile
	registerUser.Avatar = avatar
	registerUser.Nickname = nickname
	registerUser.Sex = sex
	registerUser.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	registerUser.Passwd = util.MakePasswd(plainPwd, registerUser.Salt)
	registerUser.Createat = time.Now()
	//插入用户信息
	_, err = model.DbEngine.InsertOne(&registerUser)

	return registerUser,  err
}
......
############################
//main.go
############################
......
func main() {
    http.HandleFunc("/user/register", controller.UserRegister)
}
```

首先我们使用`util.Bind(request, &user)`将用户参数绑定到user对象上，使用的是util包中的Bind函数，具体实现细节读者可以自行研究，主要模仿了Gin框架的参数绑定，可以拿来即用，非常方便。然后我们根据用户手机号搜索数据库中是否已经存在，如果不存在就插入到数据库中，返回注册成功信息，逻辑非常简单。登录逻辑更简单:
```
############################
//app/controller/user.go
############################
...
//用户登录
func UserLogin(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	mobile := request.PostForm.Get("mobile")
	plainpwd := request.PostForm.Get("passwd")

	//校验参数
	if len(mobile) == 0 || len(plainpwd) == 0 {
		util.RespFail(writer, "用户名或密码不正确")
	}

	loginUser, err := UserService.Login(mobile, plainpwd)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, loginUser, "")
	}
}
...
############################
//app/service/user.go
############################
...
func (s *UserService) Login(mobile, plainpwd string) (user model.User, err error) {
	//数据库操作
	loginUser := model.User{}
	model.DbEngine.Where("mobile = ?", mobile).Get(&loginUser)
	if loginUser.Id == 0 {
		return loginUser, errors.New("用户不存在")
	}
	//判断密码是否正确
	if !util.ValidatePasswd(plainpwd, loginUser.Salt, loginUser.Passwd) {
		return loginUser, errors.New("密码不正确")
	}
	//刷新用户登录的token值
	token := util.GenRandomStr(32)
	loginUser.Token = token
	model.DbEngine.ID(loginUser.Id).Cols("token").Update(&loginUser)

	//返回新用户信息
	return loginUser, nil
}
...
############################
//main.go
############################
......
func main() {
    http.HandleFunc("/user/login", controller.UserLogin)
}
```

实现了登录逻辑，接下来我们就到了用户首页，这里列出了用户列表，点击即可进入聊天页面。用户也可以点击下边的tab栏查看自己所在的群组，可以由此进入群组聊天页面。具体这些工作还需要读者自己开发用户列表、添加好友、创建群组、添加群组等功能，这些都是一些普通的api开发工作，我们的代码程序中也实现了，读者可以拿去修改使用，这里就不再演示了。我们再重点看一下用户鉴权这一块吧，用户鉴权是指用户点击聊天进入聊天界面时，客户端会发送一个GET请求给服务端，请求建立一条websocket长连接，服务端收到建立连接的请求之后，会对客户端请求进行校验，以确实是否建立长连接，然后将这条长连接的句柄添加到map当中(因为服务端不仅仅对一个客户端服务，可能存在千千万万个长连接)维护起来。我们下边来看具体代码实现:

```
############################
//app/controller/chat.go
############################
......
//本核心在于形成userid和Node的映射关系
type Node struct {
	Conn *websocket.Conn
	//并行转串行,
	DataQueue chan []byte
	GroupSets set.Interface
}
......
//userid和Node映射关系表
var clientMap map[int64]*Node = make(map[int64]*Node, 0)
//读写锁
var rwlocker sync.RWMutex
//实现聊天的功能
func Chat(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id := query.Get("id")
	token := query.Get("token")
	userId, _ := strconv.ParseInt(id, 10, 64)
	//校验token是否合法
	islegal := checkToken(userId, token)

	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return islegal
		},
	}).Upgrade(writer, request, nil)

	if err != nil {
		log.Println(err.Error())
		return
	}
	//获得websocket链接conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	//获取用户全部群Id
	comIds := concatService.SearchComunityIds(userId)
	for _, v := range comIds {
		node.GroupSets.Add(v)
	}

	rwlocker.Lock()
	clientMap[userId] = node
	rwlocker.Unlock()

	//开启协程处理发送逻辑
	go sendproc(node)

	//开启协程完成接收逻辑
	go recvproc(node)

	sendMsg(userId, []byte("welcome!"))
}

......

//校验token是否合法
func checkToken(userId int64, token string) bool {
	user := UserService.Find(userId)
	return user.Token == token
}

......

############################
//main.go
############################
......
func main() {
    http.HandleFunc("/chat", controller.Chat)
}
......
```

进入聊天室，客户端发起`/chat`的GET请求，服务端首先创建了一个`Node`结构体，用来存储和客户端建立起来的websocket长连接句柄，每一个句柄都有一个管道`DataQueue`，用来收发信息，`GroupSets`是客户端对应的群组信息，后边我们会提到。

```
type Node struct {
	Conn *websocket.Conn
	//并行转串行,
	DataQueue chan []byte
	GroupSets set.Interface
}
```

服务端创建了一个map，将客户端用户id和其`Node`关联起来：

```
//userid和Node映射关系表
var clientMap map[int64]*Node = make(map[int64]*Node, 0)
```

接下来是主要的用户逻辑了，服务端接收到客户端的参数之后，首先校验token是否合法，由此确定是否要升级http协议到websocket协议，建立长连接，这一步称为鉴权。

```
	//校验token是否合法
	islegal := checkToken(userId, token)

	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return islegal
		},
	}).Upgrade(writer, request, nil)
```

鉴权成功以后，服务端初始化一个`Node`,搜索该客户端用户所在的群组`id`,填充到群组的`GroupSets`属性中。然后将`Node`节点添加到`ClientMap`中维护起来，我们对`ClientMap`的操作一定要加锁，因为Go语言在并发情况下，对map的操作并不保证原子安全:

```
//获得websocket链接conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	//获取用户全部群Id
	comIds := concatService.SearchComunityIds(userId)
	for _, v := range comIds {
		node.GroupSets.Add(v)
	}

	rwlocker.Lock()
	clientMap[userId] = node
	rwlocker.Unlock()
```

服务端和客户端建立了长链接之后，会开启两个协程专门来处理客户端消息的收发工作，对于Go语言来说，维护协程的代价是很低的，所以说我们的单机程序可以很轻松的支持成千上完的用户聊天，这还是在没有优化的情况下。

```
......
//开启协程处理发送逻辑
	go sendproc(node)

	//开启协程完成接收逻辑
	go recvproc(node)

	sendMsg(userId, []byte("welcome!"))
......	
```

至此，我们的鉴权工作也已经完成了，客户端和服务端的连接已经建立好了，接下来我们就来实现具体的聊天功能吧。

### 4.4 实现单聊和群聊

实现聊天的过程中，消息体的设计至关重要，消息体设计的合理，功能拓展起来就非常的方便，后期维护、优化起来也比较简单。我们先来看一下，我们消息体的设计:

```
############################
//app/controller/chat.go
############################
type Message struct {
	Id      int64  `json:"id,omitempty" form:"id"`           //消息ID
	Userid  int64  `json:"userid,omitempty" form:"userid"`   //谁发的
	Cmd     int    `json:"cmd,omitempty" form:"cmd"`         //群聊还是私聊
	Dstid   int64  `json:"dstid,omitempty" form:"dstid"`     //对端用户ID/群ID
	Media   int    `json:"media,omitempty" form:"media"`     //消息按照什么样式展示
	Content string `json:"content,omitempty" form:"content"` //消息的内容
	Pic     string `json:"pic,omitempty" form:"pic"`         //预览图片
	Url     string `json:"url,omitempty" form:"url"`         //服务的URL
	Memo    string `json:"memo,omitempty" form:"memo"`       //简单描述
	Amount  int    `json:"amount,omitempty" form:"amount"`   //其他和数字相关的
}
```

每一条消息都有一个唯一的`id`，将来我们可以对消息持久化存储，但是我们系统中并没有做这件工作，读者可根据需要自行完成。然后是`userid`，发起消息的用户，对应的是`dstid`,要将消息发送给谁。还有一个参数非常重要，就是`cmd`,它表示是群聊还是私聊，群聊和私聊的代码处理逻辑有所区别，我们为此专门定义了一些`cmd`常量:

```
//定义命令行格式
const (
	CmdSingleMsg = 10
	CmdRoomMsg   = 11
	CmdHeart     = 0
)
```

`media`是媒体类型，我们都知道微信支持语音、视频和各种其他的文件传输，我们设置了该参数之后，读者也可以自行拓展这些功能。`content`是消息文本，是聊天中最常用的一种形式。`pic`和`url`是为图片和其他链接资源所设置的。`memo`是简介，`amount`是和数字相关的信息，比如说发红包业务有可能使用到该字段。

消息体的设计就是这样，基于此消息体，我们来看一下，服务端如何收发消息，实现单聊和群聊吧。还是从上一节说起，我们为每一个客户端长链接开启了两个协程，用于收发消息，聊天的逻辑就在这两个协程当中实现。

```
############################
//app/controller/chat.go
############################
......
//发送逻辑
func sendproc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

//接收逻辑
func recvproc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}

		dispatch(data)
		//todo对data进一步处理
		fmt.Printf("recv<=%s", data)
	}
}
......
//后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	switch msg.Cmd {
	case CmdSingleMsg:
		sendMsg(msg.Dstid, data)
	case CmdRoomMsg:
		for _, v := range clientMap {
			if v.GroupSets.Has(msg.Dstid) {
				v.DataQueue <- data
			}
		}
	case CmdHeart:
		//检测客户端的心跳
	}
}

//发送消息,发送到消息的管道
func sendMsg(userId int64, msg []byte) {
	rwlocker.RLock()
	node, ok := clientMap[userId]
	rwlocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
......
```

服务端向客户端发送消息逻辑比较简单，就是将客户端发送过来的消息，直接添加到目标用户`Node`的channel中去就好了。通过websocket的`WriteMessage`就可以实现此功能：

```
func sendproc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}
```

收发逻辑是这样的，服务端通过websocket的`ReadMessage`方法接收到用户信息，然后通过`dispatch`方法进行调度:

```
func recvproc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}

		dispatch(data)
		//todo对data进一步处理
		fmt.Printf("recv<=%s", data)
	}
}
```

`dispatch`方法所做的工作有两件：

- 解析消息体到Message中
- 根据消息类型，将消息体添加到不同用户或者用户组的channel当中

Go语言中的channel是协程间通信的强大工具, `dispatch`只要将消息添加到channel当中，发送协程就会获取到信息发送给客户端，这样就实现了聊天功能，单聊和群聊的区别只是服务端将消息发送给群组还是个人，如果发送给群组，程序会遍历整个`clientMap`, 看看哪个用户在这个群组当中，然后将消息发送。其实更好的实践是我们再维护一个群组和用户关系的Map，这样在发送群组消息的时候，取得用户信息就比遍历整个`clientMap`代价要小很多了。

```
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	switch msg.Cmd {
	case CmdSingleMsg:
		sendMsg(msg.Dstid, data)
	case CmdRoomMsg:
		for _, v := range clientMap {
			if v.GroupSets.Has(msg.Dstid) {
				v.DataQueue <- data
			}
		}
	case CmdHeart:
		//检测客户端的心跳
	}
}
......
func sendMsg(userId int64, msg []byte) {
	rwlocker.RLock()
	node, ok := clientMap[userId]
	rwlocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
```

可以看到，通过channel，我们实现用户聊天功能还是非常方便的，代码可读性很强，构建的程序也很健壮。下边是笔者本地聊天的示意图：

<center class="half">
    <img src="https://user-gold-cdn.xitu.io/2020/1/18/16fb85df3fad2b59?w=566&h=992&f=png&s=83885" width="200"/>
    <img src="https://user-gold-cdn.xitu.io/2020/1/18/16fb85e18d870051?w=562&h=1004&f=png&s=84831" width="200"/><img src="图片链接" width="200"/>
</center>

### 4.5 发送表情和图片

下边我们再来看一下聊天中经常使用到的发送表情和图片功能是如何实现的吧。其实表情也是小图片，只是和聊天中图片不同的是，表情图片比较小，可以缓存在客户端，或者直接存放到客户端代码的代码文件中（不过现在微信聊天中有的表情包都是通过网络传输的）。下边是一个聊天中返回的图标文本数据：

```
{
"dstid":1,
"cmd":10,
"userid":2,
"media":4,
"url":"/asset/plugins/doutu//emoj/2.gif"
}
```

客户端拿到url后，就加载本地的小图标。聊天中用户发送图片也是一样的原理，不过聊天中用户的图片需要先上传到服务器，然后服务端返回url，客户端再进行加载，我们的IM系统也支持此功能，我们看一下图片上传的程序:

```
############################
//app/controller/upload.go
############################
func init() {
	os.MkdirAll("./resource", os.ModePerm)
}

func FileUpload(writer http.ResponseWriter, request *http.Request) {
	UploadLocal(writer, request)
}

//将文件存储在本地/im_resource目录下
func UploadLocal(writer http.ResponseWriter, request *http.Request) {
	//获得上传源文件
	srcFile, head, err := request.FormFile("file")
	if err != nil {
		util.RespFail(writer, err.Error())
	}
	//创建一个新的文件
	suffix := ".png"
	srcFilename := head.Filename
	splitMsg := strings.Split(srcFilename, ".")
	if len(splitMsg) > 1 {
		suffix = "." + splitMsg[len(splitMsg)-1]
	}
	filetype := request.FormValue("filetype")
	if len(filetype) > 0 {
		suffix = filetype
	}
	filename := fmt.Sprintf("%d%s%s", time.Now().Unix(), util.GenRandomStr(32), suffix)
	//创建文件
	filepath := "./resource/" + filename
	dstfile, err := os.Create(filepath)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}
	//将源文件拷贝到新文件
	_, err = io.Copy(dstfile, srcFile)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}

	util.RespOk(writer, filepath, "")
}
......
############################
//main.go
############################
func main() {
    http.HandleFunc("/attach/upload", controller.FileUpload)
}
```

我们将文件存放到本地的一个磁盘文件夹下，然后发送给客户端路径，客户端通过路径加载相关的图片信息。

关于发送图片，我们虽然实现功能，但是做的太简单了，我们在接下来的章节详细的和大家探讨一下系统优化相关的方案。怎样让我们的系统在生产环境中用的更好。

## 5. 程序优化和系统架构升级方案

我们上边实现了一个功能健全的IM系统，要将该系统应用在企业的生产环境中，需要对代码和系统架构做优化，才能实现真正的高可用。本节主要从代码优化和架构升级上谈一些个人观点，能力有限不可能面面俱到，希望读者也在评论区给出更多好的建议。

### 5.1 代码优化

我们的代码没有使用框架，函数和api都写的比较简陋，虽然进行了简单的结构化，但是很多逻辑并没有解耦，所以建议大家业界比较成熟的框架对代码进行重构，Gin就是一个不错的选择。

系统程序中使用`clientMap`来存储客户端长链接信息，Go语言中对于大Map的读写要加锁，有一定的性能限制，在用户量特别大的情况下，读者可以对`clientMap`做拆分，根据用户id做hash或者采用其他的策略，也可以将这些长链接句柄存放到redis中。

上边提到图片上传的过程，有很多可以优化的地方，首先是图片压缩(微信也是这样做的)，图片资源的压缩不仅可以加快传输速度，还可以减少服务端存储的空间。另外对于图片资源来说，实际上服务端只需要存储一份数据就够了，读者可以在图片上传的时候做hash校验，如果资源文件已经存在了，就不需要再次上传了，而是直接将url返回给客户端（各大网盘厂商的妙传功能就是这样实现的）

代码还有很多优化的地方，比如我们可以将鉴权做的更好，使用wss://代替ws://，在一些安全领域，可以对消息体进行加密，在高并发领域，可以对消息体进行压缩;对Mysql连接池再做优化，将消息持久化存储到mongo，避免对数据库频繁的写入，将单条写入改为多条一块写入;为了使程序耗费更少的Cpu,降低对消息体进行Json编码的次数，一次编码，多次使用......

### 5.2 系统架构升级

我们的系统太过于简单，所在在架构升级上，有太多的工作可以做，笔者在这里只提几点比较重要的：

- 应用/资源服务分离

我们所说的资源指的是图片、视频等文件，可以选择成熟厂商的Cos，或者自己搭建文件服务器也是可以的，如果资源量比较大，用户比较广，cdn是不错的选择。

- 突破系统连接数,搭建分布式环境

对于服务器的选择，一般会选择linux，linux下一切皆文件，长链接也是一样。单机的系统连接数是有限制的，一般来说能达到10万就很不错了，所以在用户量增长到一定程序，需要搭建分布式。分布式的搭建就要优化程序，因为长链接句柄分散到不同的机器，实现消息广播和分发是首先要解决的问题，笔者这里不深入阐述了，一来是没有足够的经验，二来是解决方案有太多的细节需要探讨。搭建分布式环境所面临的问题还有：怎样更好的弹性扩容、应对突发事件等。

- 业务功能分离

我们上边将用户注册、添加好友等功能和聊天功能放到了一起，真实的业务场景中可以将它们做分离，将用户注册、添加好友、创建群组放到一台服务器上，将聊天功能放到另外的服务器上。业务的分离不仅使功能逻辑更加清晰，还能更有效的利用服务器资源。

- 减少数据库I/O,合理利用缓存

我们的系统没有将消息持久化，用户信息持久化到mysql中去。在业务当中，如果要对消息做持久化储存，就要考虑数据库I/O的优化，简单讲：合并数据库的写次数、优化数据库的读操作、合理的利用缓存。

上边是就是笔者想到的一些代码优化和架构升级的方案。

## 6.结束语

不知道大家有没有发现，使用Go搭建一个IM系统比使用其他语言要简单很多，而且具备更好的拓展性和性能(并没有吹嘘Go的意思)。在当今这个时代，5G将要普及，流量不再昂贵，IM系统已经广泛渗入到了用户日常生活中。对于程序员来说，搭建一个IM系统不再是困难的事情, 如果读者根据本文的思路，理解Websocket，Copy代码，运行程序,应该用不了半天的时间就能上手这样一个IM系统。IM系统是一个时代，从QQ、微信到现在的人工智能，都广泛应用了即时通信，围绕即时通信，又可以做更多产品布局。笔者写本文的目的就是想要帮助更多人了解IM，帮助一些开发者快速的搭建一个应用，燃起大家学习网络编程知识的兴趣，希望的读者能有所收获，能将IM系统应用到更多的产品布局中。


























