
gateway_list()
function gateway_list() {
    var success = function (obj) {
        $('#gateway-list').empty()
        $.each(obj, function (i, val) {
            var list = `<tr>
                    <td>`+ val.gateway_id + `</td>
                    <td class="text-center">`+ val.gateway_name + `</td>
                    <td class="text-center"><span id="`+ val.gateway_id + `"></span></td>
                    <td class="text-center">
                        <button type="button" onclick="get_gateway('`+val.gateway_id+`')" class="btn btn-sm btn-success">
                            <i class="icon-edit"></i> 编辑</button>
                        <button type="button" onclick="del_gateway('`+val.gateway_id+`')"  class="btn btn-sm btn-danger">
                            <i class="icon-trash"></i> 删除</button>
                    </td>
                </tr>`
            $('#gateway-list').append(list)
            $('#' + val.gateway_id).qrcode({ width: 64, height: 64, text: val.gateway_id });
        });
    }
    AjaxReq(
        "get",
        "api/device/gateway/find",
        {},
        function () { },
        success,
        ReqErr
    )
}

$('#gateway-commit').click(function () {
        add_gateway()
});

$("#add-gateway").click(function(){
    clean_gateway()
})

function get_gateway(id){
    AjaxReq(
        "get",
        "api/device/gateway/get",
        {gateway_id:id},
        function() {},
        function(msg) {
            $('#gateway-id').val(msg.gateway_id)
            $('#gateway-name').val(msg.gateway_name)
            $("#gateway-modal").modal("show")
        },
        ReqErr
    )
}

function del_gateway(id){
    AjaxReq(
        "get",
        "api/device/gateway/delete",
        {gateway_id:id},
        function() {},
        function(msg) {
            gateway_list()
            ReqSuccess()
        },
        ReqErr
    )
}

function add_gateway() {
    var id=$('#gateway-id').val()
    var name = $('#gateway-name').val()
   
    var url="api/device/gateway/create"
    if (id!=""){
        url="api/device/gateway/update"
    } 
    AjaxReq(
        "post",
        url,
        { gateway_name: name,gateway_id:id},
        function() {},
        function() {
            gateway_list()
            clean_gateway()
            $("#gateway-modal").modal("hide")
            ReqSuccess()
        },
        ReqErr
    )
}

function clean_gateway(){
    $('#gateway-id').val("")
    $('#gateway-name').val("")
}
