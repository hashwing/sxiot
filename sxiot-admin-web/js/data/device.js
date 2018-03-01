$(function () {
    device_list()
    tpl_list()
})

function device_list() {
    var success = function (obj) {
        $('#device-list').empty()
        $.each(obj, function (i, val) {
            var list = `<tr>
                    <td>`+ val.device_id + `</td>
                    <td class="text-center">`+ val.device_alias + `</td>
                    <td class="text-center">
                        <button type="button" onclick="get_device('`+val.device_id+`')" class="btn btn-sm btn-success">
                            <i class="icon-edit"></i> 编辑</button>
                        <button type="button" onclick="del_device('`+val.device_id+`')" class="btn btn-sm btn-danger">
                            <i class="icon-trash"></i> 删除</button>
                    </td>
                </tr>`
            $('#device-list').append(list)
        });
    }
    AjaxReq(
        "get",
        "api/device/find",
        {},
        function () { },
        success,
        ReqErr
    )

}



$('#device-commit').click(function () {
    add_device()
});

$('#add-device').click(function () {
    clean_device()
})

function add_device() {
    var name = $('#device-name').val()
    var brand_id = $('#brand-id').val()
    var device_unit = $('#device-unit').val()
    var id = $('#device-id').val()
    var url = "api/device/create"
    if (id != "") {
        url = "api/device/update"
    }
    AjaxReq(
        "post",
        url,
        { device_id: id, device_name: name, brand_id: brand_id, device_unit: device_unit },
        function () { },
        function () {
            device_list()
            clean_device()
            $("#device-modal").modal("hide")
            ReqSuccess()
        },
        ReqErr
    )
}

function get_device(id) {
    AjaxReq(
        "get",
        "api/device/get",
        { device_id: id },
        function () { },
        function (msg) {
            $('#device-id').val(msg.device_id)
            $('#device-name').val(msg.device_alias)
            $('#brand-id').val(msg.brand_id)
            $('#device-unit').val(msg.device_unit)
            $("#device-modal").modal("show")
        },
        ReqErr
    )
}


function del_device(id){
    AjaxReq(
        "get",
        "api/device/delete",
        {device_id:id},
        function() {},
        function() {
            device_list()
            ReqSuccess()
        },
        ReqErr
    )
}

function tpl_list() {
    $.ajax({
        type: "get",
        url: "api/device/brand/find",
        data: {},
        beforeSend: function () { },
        success: function (msg) {
            var obj = JSON.parse(msg)
            $('#brand-id').empty()
            $.each(obj, function (i, val) {
                var list = `<option value="` + val.brand_id + `">` + val.brand_name + `</option>`
                $('#brand-id').append(list)
            });
        },
        error: function (XMLHttpRequest, textStatus, errorThrown) {
        }
    });


}

function clean_device() {
    $('#device-id').val("")
    $('#device-name').val("")
    $('#brand-id').val("")
    $('#device-unit').val("")
}