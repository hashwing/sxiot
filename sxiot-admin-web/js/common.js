function AjaxReq(type, url,data, before, success, err) {
    $.ajax({
      type: type,
      url: url,
      dataType:'json',
      data: data,
      beforeSend: function (xhr) {
        xhr.setRequestHeader('Authorization', $.cookie("Auth"));
        before
      },
      success: success,
      error: function (XMLHttpRequest, textStatus, errorThrown) {
        if (XMLHttpRequest.status==401||XMLHttpRequest.status==403){
          $(top.location).attr("href","login.html")
        }
        err
      }
    })
  }
  
  function ReqErr(){
    new $.zui.Messager('提示消息：请求失败', {
        type: 'danger' 
    }).show();
  }
  
  function ReqSuccess(){
    new $.zui.Messager('提示消息：操作成功', {
      type: 'success' 
  }).show();
  }