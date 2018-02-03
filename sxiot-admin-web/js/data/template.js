
tpl_list()
function tpl_list(){
    $.ajax({
        type: "get",
        url: "api/device/brand/find",
        data: {},
        beforeSend: function() {},
        success: function(msg) {
            var obj= JSON.parse(msg)
            $('#tpl-list').empty()
            $.each(obj,function(i,val){
                var list= `<tr>
                    <td>`+val.brand_name+`</td>
                    <td class="text-center">`+val.brand_type+`</td>
                    <td class="text-center">
                        <button type="button" class="btn btn-sm btn-success">
                            <i class="icon-edit"></i> 编辑</button>
                        <button type="button" class="btn btn-sm btn-danger">
                            <i class="icon-trash"></i> 删除</button>
                    </td>
                </tr>`
                $('#tpl-list').append(list)
            });   
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
        }
    });

    
}

$('#commit').click(function(){
    if ($('#tpl-id').val()==""){
        add_tpl()
    }
});

function add_tpl(){
    var name=$('#tpl-name').val()
    var type=$('#tpl-mark').val()
    var metadata=$('#tpl-metadata').val()
    $.ajax({
        type: "post",
        url: "api/device/brand/create",
        data: {brand_name:name,brand_type:type,brand_metadata:metadata},
        beforeSend: function() {},
        success: function(msg) {
            new $.zui.Messager("添加成功", {
                type: 'success' 
                }).show();  
                tpl_list()   
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            new $.zui.Messager("添加失败", {
                type: 'danger' 
                }).show();
        }
    });
}
