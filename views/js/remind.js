
function ShowDanger() {
    var s="";
    box.style.display = 'flex';
    hidden.style.display = 'block';
    // window.alert("SHOW")
    $.ajax({
        url: '/remind/getun',
        type: 'get',
        dataType: 'json',
        success: function (res) {
            console.log(res)
            // document.getElementById("un_name").innerHTML=res.msg;
            for(var i=0;i<res.msg.length;i++){
                if(i==res.msg.length-1)
                {
                    s=s+res.msg[i];
                }else{
                    s=s+res.msg[i]+",";
                }
            }
            console.log(s)
            $('#un_name').html(s)
        }
    })
    ShowDangerto()


}


function ShowDangerto(){
    var s="";
    box.style.display = 'flex';
    hidden.style.display = 'block';
    // window.alert("SHOW")
    $.ajax({
        url: '/remind/getdanger',
        type: 'get',
        dataType: 'json',
        success: function (res) {
            console.log(res)
            // document.getElementById("un_name").innerHTML=res.msg;
            for(var i=0;i<res.msg.length;i++){
                if(i==res.msg.length-1)
                {
                    s=s+res.msg[i];
                }else{
                    s=s+res.msg[i]+",";
                }
            }
            console.log(s)
            $('#danger_name').html(s)
        }
    })
}


function my$(id) {
    return document.getElementById(id);
}

// 获取鼠标在页面的位置，处理浏览器兼容性
function getPage(e) {
    var pageX = e.pageX || e.clientX + getScroll().scrollLeft;
    var pageY = e.pageY || e.clientY + getScroll().scrollTop;
    return {
        pageX: pageX,
        pageY: pageY
    }
}


