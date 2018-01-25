
device_list()
tpl_list()
function device_list(){
    $.ajax({
        type: "get",
        url: "api/device/find",
        data: {},
        beforeSend: function() {},
        success: function(msg) {
            var obj= JSON.parse(msg)
            $('#device-list').empty()
            $.each(obj,function(i,val){
                var list= `<tr>
                    <td>`+val.device_id+`</td>
                    <td class="text-center">`+val.device_alias+`</td>
                    <td class="text-center">
                        <button type="button" class="btn btn-sm btn-success">
                            <i class="icon-edit"></i> 编辑</button>
                        <button type="button" class="btn btn-sm btn-danger">
                            <i class="icon-trash"></i> 删除</button>
                    </td>
                </tr>`
                $('#device-list').append(list)
            });   
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
        }
    });

    
}

$('#commit').click(function(){
    if ($('#device-id').val()==""){
        add_device()
    }
});

function add_device(){
    var name=$('#device-name').val()
    var brand_id=$('#brand-id').val()
    var device_unit=$('#device-unit').val()
    alert(device_unit)
    $.ajax({
        type: "post",
        url: "api/device/create",
        data: {device_name:name,brand_id:brand_id,device_unit:device_unit},
        beforeSend: function() {},
        success: function(msg) {
            new $.zui.Messager("添加成功", {
                type: 'success' 
                }).show();  
                device_list()   
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            new $.zui.Messager("添加失败", {
                type: 'danger' 
                }).show();
        }
    });
}

function tpl_list(){
    $.ajax({
        type: "get",
        url: "api/device/brand/find",
        data: {},
        beforeSend: function() {},
        success: function(msg) {
            var obj= JSON.parse(msg)
            $('#brand-id').empty()
            $.each(obj,function(i,val){
                var list= `<option value="`+val.brand_id+`">`+val.brand_name+`</option>`
                $('#brand-id').append(list)
            });   
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
        }
    });

    
}