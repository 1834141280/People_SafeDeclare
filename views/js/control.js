$(document).ready(function () {




    $('#tianbao').click(function () {
        $.ajax({
            type: 'get',
            url: '/control/write',
            dataType: 'json',
            success: function (res) {
                if (res.msg === "true") {
                    window.alert("今日已填报")

                } else {
                    window.location.replace("/declare")
                }
            }
        })


    })




    $('#tongji').click(function () {
        window.location.replace("/show")
    })






    $("#search").click(function () {
        let form = {
            id: $("#chaxun").val()
        }
        $.ajax({
            url: '/control/search',
            type: 'post',
            data: form,
            dataType: 'json',
            success: function (res) {
                console.log(res.msg)
                $('#name').val(res.msg.name)
                $('#idnumber').val(res.msg.idnumber)
                $('#username').val(res.msg.id)
                $('#phonenumber').val(res.msg.name)
                $('#city').val(res.msg.city)
                $('#academy').val(res.msg.academy)
                $('#high_region').val(res.msg.high_region)
                $('#isolation').val(res.msg.isolation)
                $('#typical_symptoms').val(res.msg.typical_symptoms)
                $('#fever').val(res.msg.fever)
                $('#other').val(res.msg.other)
                $('#temperature_morning').val(res.msg.temperature_morning)
                $('#temperature_moon').val(res.msg.temperature_moon)
                $('#temperature_aftermoon').val(res.msg.temperature_aftermoon)
                $('#date').val(res.msg.date)
                $('#write').val(res.msg.write)
                // window.alert("1")
                // $('#editModal').modal('hide');
            }
        }

        )
    })



    $("#increase").click(function () {
        let id = $("#tianjia").val()
        $('#usernamein').val(id)
    })


    $("#edit").click(function () {
        let id = $("#xiugai").val()
        $('#usernameed').val(id)
    })


    $("#delete").click(function () {
        let form = {
            id: $("#shanchu").val()
        }
        console.log(form.id)
        $.ajax({
            url: '/control/delete',
            type: 'post',
            data: form,
            dataType: 'json',
            success: function (res) {
                console.log(res.msg)
                if (res.msg == "1") {
                    window.alert("删除成功")
                } else {
                    window.alert("删除失败")
                }
            }
        })





    })

})

function IncreaseData() {
    let form = {
        id: $("#usernamein").val(),
        idnumber: $("#idnumberin").val(),
        name: $("#namein").val(),
        phone_number: $("#phonenumberin").val(),
        academy: $("#academyin").val(),
        city: $("#cityin").val(),
        high_region: $("#high_regionin").val(),
        isolation: $("#isolationin").val(),
        typical_symptoms: $("#typical_symptomsin").val(),
        fever: $("#feverin").val(),
        other: $("#otherin").val(),
        temperature_morning: $("#temperature_morningin").val(),
        temperature_moon: $("#temperature_moonin").val(),
        temperature_aftermoon: $("#temperature_aftermoonin").val(),
        date: $("#datein").val(),
        write: $("#writein").val(),
    }
    let json = JSON.stringify(form)
    $.ajax({

        url: '/control/increase',
        type: 'post',
        data: json,
        dataType: 'json',
        success: function (res) {
            $('#username').val(tianjia)
            console.log(res.msg)
            if (res.msg == "1") {
                window.alert("添加成功")
            } else {
                window.alert("添加失败")
            }
        }
    }

    )
}


function EditData() {
    let form = {
        id: $("#usernameed").val(),
        idnumber: $("#idnumbered").val(),
        name: $("#nameed").val(),
        phone_number: $("#phonenumbered").val(),
        academy: $("#academyed").val(),
        city: $("#cityed").val(),
        high_region: $("#high_regioned").val(),
        isolation: $("#isolationed").val(),
        typical_symptoms: $("#typical_symptomsed").val(),
        fever: $("#fevered").val(),
        other: $("#othered").val(),
        temperature_morning: $("#temperature_morninged").val(),
        temperature_moon: $("#temperature_mooned").val(),
        temperature_aftermoon: $("#temperature_aftermooned").val(),
        date: $("#dateed").val(),
        write: $("#writeed").val(),
    }
    let json = JSON.stringify(form)
    $.ajax({

        url: '/control/edit',
        type: 'post',
        data: json,
        dataType: 'json',
        success: function (res) {
            $('#username').val(tianjia)
            console.log(res.msg)
            if (res.msg == "1") {
                window.alert("修改成功")
            } else {
                window.alert("修改失败")
            }
        }
    })
}

