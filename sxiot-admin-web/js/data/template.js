
tpl_list()
function tpl_list(){
    var success =function(obj) {
        $('#tpl-list').empty()
        $.each(obj,function(i,val){
            var list= `<tr>
                <td>`+val.brand_name+`</td>
                <td class="text-center">`+val.brand_type+`</td>
                <td class="text-center">
                    <button type="button" onclick="get_tpl('`+val.brand_id+`')" class="btn btn-sm btn-success">
                        <i class="icon-edit"></i> 编辑</button>
                    <button type="button" onclick="del_tpl('`+val.brand_id+`')" class="btn btn-sm btn-danger">
                        <i class="icon-trash"></i> 删除</button>
                </td>
            </tr>`
            $('#tpl-list').append(list)
        });   
    }
    AjaxReq(
        "get",
        "api/device/brand/find",
        {},
        function() {},
        success,
        ReqErr
    )
}

$('#tpl-commit').click(function(){
    add_tpl()
});

function add_tpl(){
    var id=$('#tpl-id').val()
    var name=$('#tpl-name').val()
    var type=$('#tpl-mark').val()
    var metadata=$('#tpl-metadata').val()
    var url="api/device/brand/create"
    if (id!=""){
        url="api/device/brand/update"
    }
    AjaxReq(
        "post",
        url,
        {brand_id:id,brand_name:name,brand_type:type,brand_metadata:metadata},
        function() {},
        function() {
            tpl_list()
            clean_tpl()
            $("#tpl-modal").modal("hide")
            ReqSuccess()
        },
        ReqErr
    )
}

function get_tpl(id){
    AjaxReq(
        "get",
        "api/device/brand/get",
        {brand_id:id},
        function() {},
        function(msg) {
            $('#tpl-id').val(msg.brand_id)
            $('#tpl-name').val(msg.brand_name)
            $('#tpl-mark').val(msg.brand_type)
            $('#tpl-metadata').val(msg.brand_metadata)
            $("#tpl-modal").modal("show")
        },
        ReqErr
    )
}

function del_tpl(id){
    AjaxReq(
        "get",
        "api/device/brand/delete",
        {brand_id:id},
        function() {},
        function() {
            tpl_list()
            ReqSuccess()
        },
        ReqErr
    )
}

function clean_tpl(){
    $('#tpl-id').val("")
    $('#tpl-name').val("")
    $('#tpl-mark').val("")
    $('#tpl-metadata').val("")
}
