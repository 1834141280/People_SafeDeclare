$(document).ready(function () {
    $('#submitsign').click(function () {
        let user = {
            username: $('#usernamesign').val(),
            password: $('#passwordsign').val()
        }
        // console.log(user)
        if (user.username.length > 12 || user.username.length < 6 || user.password.length < 6 || user.password.length > 12) {
            $('#usernamesign').val("")
            window.alert("用户名或密码格式错误，请重新输入")
        } else {
            let json = JSON.stringify(user)
            console.log(json)
            $.ajax({
                type: 'post',
                url: '/index/sign',
                dataType: 'json',
                data: json,
                success: function (res) {
                    console.log(res.msg)
                    if (res.msg === 0) {
                        window.location.replace("/remind")
                    } else if (res.msg === 1) {
                        window.location.replace("/declare")
                    } else if (res.msg === 2) {
                        $('#usernamesign').val("")
                        $('#passwordsign').val("")
                        window.alert("用户名或密码错误")
                    } else {
                        $('#usernamesign').val("")
                        $('#passwordsign').val("")
                        window.alert("服务器错误，请稍后在试")
                    }
                }
            })
        }
    })
    $('#submitregister').click(function () {
        let user = {
            identity: $('#identityregister').val(),
            username: $('#usernameregister').val(),
            password: $('#passwordregister').val(),
            useremail: $('#useremailregister').val()
        }
        console.log(user)
        if (user.identity == 0) {
            window.alert("请选择身份")
        }
        else if (user.username.length > 12 || user.username.length < 6 ||
            user.username.password > 12 || user.username.password < 6) {
            $('#usernameregister').val("")
            $('#passwordregister').val("")
            $('#useremailregister').val("")
            window.alert("格式错误，请重新输入")
        } else {
            let json = JSON.stringify(user)
            console.log(json)
            $.ajax({
                type: 'post',
                url: '/index/register',
                dataType: 'json',
                data: json,
                success: function (res) {
                    if (res.msg === 1) {
                        console.log(res.msg)
                        if (user.identity === "系统管理员") {
                            window.alert("注册成功")
                            window.location.replace("/control")
                        } else {
                            window.alert("注册成功")
                            window.location.replace("/declare")
                        }
                    } else {
                        $('#usernameregister').val("")
                        $('#passwordregister').val("")
                        $('#useremailregister').val("")
                        window.alert("用户名重复，请换一个")
                    }
                }
            })
        }
    })
})