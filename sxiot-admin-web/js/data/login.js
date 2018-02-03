$('#commit').click(function(){
    var account =$('#account').val()
    var pwd=$('#pwd').val()
    $.ajax({
        type: "post",
        url: "api/login",
        data: {account:account,password:pwd},
        beforeSend: function() {},
        success: function(msg) {
            var obj= JSON.parse(msg)
            $.cookie("Auth",obj.token,{path:'/'})
            $.cookie("Auth",obj.token)
            $(location).attr("href","index.html")
            new $.zui.Messager("登录成功", {
                type: 'success' 
                }).show();    
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            new $.zui.Messager("登录失败", {
                type: 'danger' 
                }).show();
        }
    });
})