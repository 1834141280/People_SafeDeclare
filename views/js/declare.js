$(document).ready(function () {
    $('#submitform').click(function () {

        let form = {
            name: $("#subname").val(),
            idnumber: $("#subidnumber").val(),
            id: $("#subid").val(),
            phone_number: $("#subphone_number").val(),
            academy: $("#subacademy").val(),
            city: $("#subcity").val(),
            high_region: $("#subhigh_region").val(),
            isolation: $("#subisolation").val(),
            typical_symptoms: $("#subtypical_symptoms").val(),
            fever: $("#subfever").val(),
            other: $("#subother").val(),
            temperature_morning: $("#subtemperature_morning").val(),
            temperature_moon: $("#subtemperature_moon").val(),
            temperature_aftermoon: $("#subtemperature_aftermoon").val(),
        }
        console.log(form)


        if (form.name.length == 0) {
            $('#subname').val("")
            window.alert("姓名填写错误，请重新输入")
        } else if (form.idnumber.length != 18) {
            $('#subidnumber').val("")
            window.alert("身份证号填写错误，请重新输入")
        } else if (form.id.length < 6 || form.id.length > 12) {
            $('#subid').val("")
            window.alert("学工号填写错误，请重新输入")
        } else if (form.phone_number.length != 11) {
            $('#subphone_number').val("")
            window.alert("手机号填写错误，请重新输入")
        }
        else if (form.academy == 0) {
            window.alert("请填写单位/院系")
        }
        else if (form.city == 0) {
            $('#subcity').val("")
            window.alert("城市填写错误，请重新输入")
        }
        else if (form.high_region == 0) {
            window.alert("请填写七天内受否有高低风险旅居史、接触史")
        }
        else if (form.isolation == 0) {
            window.alert("请填写是否隔离观察")
        }
        else if (form.typical_symptoms == 0) {
            window.alert("请填写是否身体有疑似典型症状")
        }
        else if (form.fever.length == 0) {
            window.alert("请填写是否发热")
        }
        else if (form.temperature_morning.length == 0 || form.temperature_morning.length > 4) {
            $('#subtemperature_morning').val("")
            window.alert("体温（早）填写错误，请重新输入")
        }
        else if (form.temperature_moon.length == 0 || form.temperature_moon.length > 4) {
            $('#subtemperature_moon').val("")
            window.alert("体温（午）填写错误，请重新输入")
        }
        else if (form.temperature_aftermoon.length == 0 || form.temperature_aftermoon.length > 4) {
            $('#subtemperature_aftermoon').val("")
            window.alert("体温（晚）填写错误，请重新输入")
        } else {
            let json = JSON.stringify(form)
            console.log(json)

            $.ajax({
                type: 'post',
                url: '/declare/submit',
                dataType: 'json',
                data: json,
                success: function (res) {
                    if (res.msg == 0) {
                        window.alert("提交成功")
                    } else if (res.msg == 1) {
                        window.alert("今日已提交，无需重复提交")
                    } else {
                        window.alert("服务器错误，请稍后在试")
                    }
                }

            })
        }


    })
})