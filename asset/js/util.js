String.prototype.startWith=function(str){  
		        if(str==null||str==""||this.length==0||str.length>this.length)  
		          return false;  
		        if(this.substr(0,str.length)==str)  
		          return true;  
		        else  
		          return false;  
		        return true;  
		    }  
String.prototype.endWith=function(str){  
		        if(str==null||str==""||this.length==0||str.length>this.length)  
		          return false;  
		        if(this.substring(this.length-str.length)==str)  
		          return true;  
		        else  
		          return false;  
		        return true;  
}

//日期格式化
Date.prototype.format = function(format){ 
	if(!format){
		format = 'yyyy-MM-dd';// 默认1997-01-01这样的格式
	}
	var o = { 
		"M+" : this.getMonth()+1, // month
		"d+" : this.getDate(), // day
		"h+" : this.getHours(), // hour
		"m+" : this.getMinutes(), // minute
		"s+" : this.getSeconds(), // second
		"q+" : Math.floor((this.getMonth()+3)/3), // quarter
		"S" : this.getMilliseconds() // millisecond
	} 

	if(/(y+)/.test(format)) { 
		format = format.replace(RegExp.$1, (this.getFullYear()+"").substr(4 - RegExp.$1.length)); 
	} 

	for(var k in o) { 
		if(new RegExp("("+ k +")").test(format)) { 
			format = format.replace(RegExp.$1, RegExp.$1.length==1 ? o[k] : ("00"+ o[k]).substr((""+ o[k]).length)); 
		} 
	} 
	return format; 
} 


var config={
	serverUrl:"//"+location.host
}

function Core(){
	
}
Core.prototype.numformat=function(){
	  var num = (num || 0).toString(), result = '';
    var suffix="";
    if(num.indexOf(".")>-1){
    	var t = num.split(".");
    	num=t[0];
    	suffix = "." + t[1];
    	
    }
   
    while (num.length > 3) {
        result = ',' + num.slice(-3) + result;
        num = num.slice(0, num.length - 3);
    }
    if (num) { result = num +""+ result; }
  
    return result+suffix;
}
Core.prototype.token=function(d){
		d.createAt = new Date();
		return this.data("token",d);
}

Core.prototype.error=function(d){
    alert(d)
}
Core.prototype.alert=function(d){
    alert(d)
}
	//用于存储信息
Core.prototype.data=function(k,d){
		//console.log("data",k,d)
		if(typeof d=="undefined"){
			var o = localStorage.getItem(k);
			if(o==null){
				return null;
			}else{
				o = JSON.parse(o)
				return o[k]
			}
			
			
		}else if(null==d){
			return localStorage.removeItem(k)
		}else{
			var o = {}
			o[k] = d;
			return localStorage.setItem(k,JSON.stringify(o))
		}
	}
//用于存储信息
Core.prototype.api=function(uri){
    if(uri.startWith("/")){
        return config.serverUrl+uri
    }else{
        return config.serverUrl+"/"+uri
    }
}
Core.prototype.post=function(uri,data,fn){
	var url = this.api(uri)
    return new Promise(function (resolve, reject) {
        var xhr = new XMLHttpRequest();
        xhr.open("POST",url, true);
        // 添加http头，发送信息至服务器时内容编码类型
        xhr.setRequestHeader(
        	"Content-Type",
			"application/x-www-form-urlencoded"
		);
        xhr.onreadystatechange = function() {
            if (xhr.readyState == 4 && (xhr.status == 200 || xhr.status == 304)) {
                resolve(JSON.parse(xhr.responseText));
            }
        };
        xhr.onerror = function(){
        	reject({"code":-1,"msg":"服务器繁忙"})
		}
        var _data=[];
        for(var i in data){
            _data.push( i +"=" + encodeURI(data[i]));
        }
    	xhr.send(_data.join("&"));

    })
}
Core.prototype.uploadfile=function(uri,dom){
    var url = this.api(uri)
    return new Promise(function (resolve, reject) {
        var xhr = new XMLHttpRequest();
        xhr.open("POST",url, true);
        xhr.onreadystatechange = function() {
            if (xhr.readyState == 4 && (xhr.status == 200 || xhr.status == 304)) {
                resolve(JSON.parse(xhr.responseText));
            }
        };
        xhr.onerror = function(){
            reject({"code":-1,"msg":"服务器繁忙"})
        }

        var formdata = new FormData();
        formdata.append("file",dom.files[0])
        xhr.send(formdata);
    })


}

Core.prototype.uploadmp3=function(uri,blob){

    var url = this.api(uri)
    return new Promise(function (resolve, reject) {
        var xhr = new XMLHttpRequest();
        xhr.open("POST",url, true);
        xhr.onreadystatechange = function() {
            if (xhr.readyState == 4 && (xhr.status == 200 || xhr.status == 304)) {
                resolve(JSON.parse(xhr.responseText));
            }
        };
        xhr.onerror = function(){
            reject({"code":-1,"msg":"服务器繁忙"})
        }

        var formdata = new FormData();
        formdata.append("filetype",".mp3");
        formdata.append("file",blob)
        xhr.send(formdata);
    })

}

Core.prototype.parseUri = function(url){
		if(typeof url=="undefined"){
			url= location.href;
		}
		var query = url.substr(url.indexOf("?"));
		query=query.substr(1);
	    var reg = /([^=&\s]+)[=\s]*([^=&\s]*)/g;
	    var obj = {};
	    while(reg.exec(query)){
	        obj[RegExp.$1] = decodeURI(RegExp.$2);
	    }
	    return obj;
	}
	Core.prototype.parseQuery = function(name){
			
            var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
            var r = window.location.search.substr(1).match(reg);  //匹配目标参数
            if (r != null) return decodeURI(unescape(r[2])); 
            return null; //返回参数值
	}

Core.prototype.isemail = function(email){
    return /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+[\.][a-zA-Z0-9_-]+$/.test(email)
}
Core.prototype.ismobile = function(mobile){
    return /^[1][34578][0-9]{9}$/.test(mobile)
}

Core.prototype.test = function(reg,data){
    var reg = new RegExp(reg); //构造一个含有目标参数的正则表达式对象
    return reg.test(data)
}
window.util = new Core();


	
