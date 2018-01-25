
gateway_list()
function gateway_list(){
    $.ajax({
        type: "get",
        url: "api/device/gateway/find",
        data: {},
        beforeSend: function() {},
        success: function(msg) {
            var obj= JSON.parse(msg)
            $('#gateway-list').empty()
            $.each(obj,function(i,val){
                var list= `<tr>
                    <td>`+val.gateway_id+`</td>
                    <td class="text-center">`+val.gateway_name+`</td>
                    <td class="text-center"><span id="`+val.gateway_id+`"></span></td>
                    <td class="text-center">
                        <button type="button" class="btn btn-sm btn-success">
                            <i class="icon-edit"></i> 编辑</button>
                        <button type="button" class="btn btn-sm btn-danger">
                            <i class="icon-trash"></i> 删除</button>
                    </td>
                </tr>`
                $('#gateway-list').append(list)
                $('#'+val.gateway_id).qrcode({width: 64,height: 64,text: val.gateway_id});
            });   
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
        }
    });

    
}

$('#commit').click(function(){
    if ($('#gateway-id').val()==""){
        add_gateway()
    }
});

function add_gateway(){
    var name=$('#gateway-name').val()
    $.ajax({
        type: "post",
        url: "api/device/gateway/create",
        data: {gateway_name:name},
        beforeSend: function() {},
        success: function(msg) {
            new $.zui.Messager("添加成功", {
                type: 'success' 
                }).show();  
                gateway_list()   
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            new $.zui.Messager("添加失败", {
                type: 'danger' 
                }).show();
        }
    });
}
