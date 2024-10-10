function Showdata() {
    $('#table').bootstrapTable({
        url: "/show/information",                           // URL"
        method: "get",                                        // 请求类型
        contentType: "application/x-www-form-urlencoded",      // post请求必须要有，否则后台接受不到参数
        datatype:"json",
        sidePagination: "server",                              // 设置在服务端还是客户端分页
        showRefresh: false,                                    // 是否刷新按钮
        sortStable: true,                                      // 是否支持排序
        cache: false,                                          // 是否使用缓存
        pagination: true,                                      // 是否显示分页
        search: false,                                         // 是否有搜索框
        clickToSelect: true,                                   // 是否点击选中行
        pageNumber: 1,                                         // 首页页码，默认为1
        pageSize: 10,                                           // 页面数据条数
        pageList: [5, 10, 20, 30],


        rowStyle: function (row, index) {
            //这里有5个取值代表5中颜色['active', 'success', 'info', 'warning', 'danger'];
            var strclass = "";
            if (row.write == "已填报") {
                strclass = 'success';//还有一个active
            }
            else if (row.write == "未填报") {
                strclass = 'danger';
            }
            else {
                return {};
            }
            return { classes: strclass }
        },



        queryParamsType: "",
        queryParams: function (params) {
            return {
                pageSize: params.pageSize,                     // 每页记录条数
                pageNumber: params.pageNumber,                 // 当前页索引

            };
        },
        columns: [{
            field: 'name',
            title: '姓名',
            align: "center",
            halign: "center",
            valign: 'middle',
        },
        {
            field: 'idnumber',
            title: '身份证号',
            align: "center",
            halign: "center",
            valign: 'middle'
        },
        {
            field: 'id',
            title: '用户名',
            align: "center",
            halign: "center",
            valign: 'middle',
        },
        {
            field: 'phone_number',
            title: '手机号',
            align: "center",
            halign: "center",
            valign: 'middle',
        },
        {
            field: 'city',
            title: '城市',
            align: "center",
            halign: "center",
            valign: 'middle'
        },
        {
            field: 'academy',
            title: '社区选择',
            align: "center",
            halign: "center",
            valign: 'middle'
        },
        {
            field: 'high_region',
            title: '七天内受否有高低风险旅居史、接触史',
            align: "center",
            halign: "center",
            valign: 'middle'
        },
        {
            field: 'isolation',
            title: '是否隔离观察',
            align: "center",
            halign: "center",
            valign: 'middle'
        },
        {
            field: 'typical_symptoms',
            title: '是否身体有疑似典型症状',
            align: "center",
            halign: "center",
            valign: 'middle'
        },
        {
            field: 'fever',
            title: '是否是否发热',
            align: "center",
            halign: "center",
            valign: 'middle'
        },
        {
            field: 'other',
            title: '其他信息',
            align: "center",
            halign: "center",
            valign: 'middle'
        },
        {
            field: 'temperature_morning',
            title: '体温（早）',
            align: "center",
            halign: "center",
            valign: 'middle'
        },
        {
            field: 'temperature_moon',
            title: '体温（午）',
            align: "center",
            halign: "center",
            valign: 'middle'
        },
        {
            field: 'temperature_aftermoon',
            title: '体温（晚）',
            align: "center",
            halign: "center",
            valign: 'middle'
        },
        {
            field: 'date',
            title: '日期',
            align: "center",
            halign: "center",
            valign: 'middle'
        },
        {
            field: 'write',
            title: '今日是否填报',
            align: 'center',
            valign: 'middle',
            width: 200,
            // events: {
            //     'click #edit': function (e, value, row, index) {
            //         $('#coursenumber').val(row.coursenumber)
            //         $('#coursename').val(row.coursename)
            //         $('#halfyear').val(row.halfyear)
            //         $('#major').val(row.major)
            //     },
            //     'click #delete': function (e, value, row, index) {
            //         deleteInfo(row.coursenumber, row.filename);
            //     }
            // },
            // formatter : function (value, row, index) {
            //     if (row['write'] === 1) {
            //         return '已填报';
            //     }
            //     if (row['write'] === 0) {
            //         return '未填报';
            //     }
            //     return value;
            // }
        }],
        onLoadSuccess: function (res) {
            console.log("加载成功");
            console.log(res.row);
        },
        onLoadError: function () {
            console.error("加载数据失败");
        },
        onDblClickRow: function (row, $element) {
            var id = row.ID;
            EditViewById(id, 'view');
        },

    })
}