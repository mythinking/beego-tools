/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-20 21:53:28
 * @LastEditTime: 2019-08-21 17:20:23
 * @LastEditors: Please set LastEditors
 */
$(function () {
    $('#myTabs a').click(function (e) {
        e.preventDefault()
        if ($(this).prop('hash') != '#tmall') {
            $('#myModal-text').text('此功能暂未开放！！！');
            $('#myModal').modal('toggle');
            return
        }
        $(this).tab('show')
    });
    $('#myTabs1 a').click(function (e) {
        e.preventDefault()
        $(this).tab('show')
    });
    $("form").submit(function(){
        var storename = $('form #storename').val();
        var cookie = $('form #cookie').val();
        var title = $('form input[name="title"]:checked').length;
        var re = new RegExp("^[a-zA-Z0-9]+$");

        if (!storename || !re.test(storename)) {
            $('#myModal-text').text('请输入合法的店铺名，如uniqlo');
        } else if (!cookie || cookie.length < 10) {
            $('#myModal-text').text('请输入合法的cookie登录信息, 打开【天猫商城】登录后请求店铺，F12复制请求中[Request Headers][cookie]');
        } else if (title < 1) {
            $('#myModal-text').text('请至少选择一列');
        } else {
            createJob();
            return false;
        }

        $('#myModal').modal('toggle');
        return false;
    });
    
    function alertDialog(txt) {
        $('#myModal-text').text(txt);
        $('#myModal').modal('toggle');
    }

    var ilog;
    function createJob() {
        var title = new Array();
        $('form input[name="title"]:checked').each(function(){
            title.push($(this).val());
        });
        $('#send').text($('#send').data('doing')).attr('disabled', 'disabled');
        $.post("/v1/crawlers/eshop",
        {
            storename:$('form #storename').val(),
            cookie:$('form #cookie').val(),
            title:title.join(','),
        },
        function(data,status){
            if (data) {
                $('#send').text($('#send').data('start')).removeAttr('disabled');
                if (data.Code !== 0) {
                    alertDialog(data.Msg);
                } else {
                    setTimeout(() => {
                        getLog()
                    }, 100);
                    //操作成功
                    $('#download').removeClass('hide');
                    $('#download').append('<p>'+(new Date()).toLocaleString()+'&nbsp;&nbsp;&nbsp;<a href="'+data.Data+'">'+data.Data+'</a></p>');
                    alertDialog("任务执行成功, 请在执行结果里面下载最新文件！！！");
                } 
            } 
            console.log(data, status)
        });
        clearInterval(ilog);
        ilog = setInterval(function(){
            getLog()
        }, 1000);
    }
    function getLog(line) {
        $.get("/v1/crawlers/log", {},
        function(data, status){
            console.log(data);
            if (data) {
                $('#log').removeClass('hide');
                if ($('#log').text() == data.Data) {
                    clearInterval(ilog);
                } else {
                    $('#log').text(data.Data)
                    var ele = document.getElementById('log');
                    ele.scrollTop = ele.scrollHeight;
                }
            }
        });
    }
});