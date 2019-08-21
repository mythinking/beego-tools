<!--
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-20 17:00:57
 * @LastEditTime: 2019-08-21 18:38:13
 * @LastEditors: Please set LastEditors
 -->
<!DOCTYPE html>

<html>
<head>
  <title>电商爬虫系统</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

  <style type="text/css">
    .tab-content {
      margin-top: 50px;
    }
  </style>
</head>

<body>

<div class="container">
  <header>
    <div class="page-header">
      <h1>电商爬虫系统 <small>欢迎使用</small> 
        <a href="/" class="btn btn-link btn-lg" title="返回首页">
        <span class="glyphicon glyphicon-home" aria-hidden="true"></span>
        </a>
      </h1>
    </div>
    <ul class="nav nav-pills" id="myTabs">
      <li role="presentation" class="active"><a href="#tmall">天猫</a></li>
      <li role="presentation"><a href="#">淘宝</a></li>
      <li role="presentation"><a href="#">京东</a></li>
      <li role="presentation"><a href="#">苏宁</a></li>
      <li role="presentation"><a href="#">网易考拉</a></li>
      <li role="presentation"><a href="#">苏宁</a></li>
      <li role="presentation"><a href="#">拼多多</a></li>
    </ul>
  </header>

  <div class="tab-content" style="margin-bottom: 50px">
    <div class="tab-pane fade in active" id="tmall">
      <div class="row">
        <div class="col-sm-12">
          <form class="form-horizontal">
            <div class="form-group">
              <label for="storename" class="col-sm-2 control-label">店铺地址</label>
              <div class="col-sm-10">
                <input type="text" class="form-control" id="storename" placeholder="如优衣库的为:uniqlo">
                <span id="helpBlock2" class="help-block">店铺链接中的前缀,如优衣库的为:uniqlo.</span>
              </div>
            </div>
            <div class="form-group">
              <label for="cookie" class="col-sm-2 control-label">导出列</label>
              <div class="col-sm-10">
                <div class="checkbox-inline">
                  <label>
                    <input type="checkbox" value="url" name="title" checked> 链接
                  </label>
                </div>
                <div class="checkbox-inline">
                  <label>
                    <input type="checkbox" value="title" name="title" > 标题
                  </label>
                </div>
                <div class="checkbox-inline">
                  <label>
                    <input type="checkbox" value="price" name="title" > 价格
                  </label>
                </div>
                <div class="checkbox-inline">
                  <label>
                    <input type="checkbox" value="item_id" name="title" > item_id
                  </label>
                </div>
              </div>
            </div>
            <div class="form-group">
              <label for="cookie" class="col-sm-2 control-label">登录信息</label>
              <div class="col-sm-10">
                <textarea class="form-control" rows="5" id="cookie"></textarea>
                <span id="helpBlock2" class="help-block">打开【天猫商城】登录后请求店铺，F12复制请求中[Request Headers->cookie]的值.</span>
              </div>
            </div>
            <div class="form-group">
              <div class="col-sm-offset-2 col-sm-2">
                <button type="submit" class="btn btn-info" id="send" data-start="启动机器" data-doing="处理中...">启动机器</button>
              </div>
              <div class="col-sm-offset-2">
                <button type="reset" class="btn btn-danger">擦干泪再来</button>
              </div>
            </div>
          </form>
        </div>
      </div>
      <!-- Small modal -->
      <div id="myModal" class="modal fade bs-example-modal-sm" tabindex="-1" role="dialog" aria-labelledby="mySmallModalLabel">
        <div class="modal-dialog modal-sm" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
              <h4 class="modal-title" id="myModalLabel">提示</h4>
            </div>
            <div class="modal-body">
              <p id="myModal-text"></p>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-info" data-dismiss="modal">知道了</button>
            </div>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-sm-12">
          <ul class="nav nav-tabs" id="myTabs1">
              <li role="presentation" class="active"><a href="#clog">日志</a></li>
            <li role="presentation"><a href="#cresult">执行结果</a></li>
          </ul>
        </div>
      </div>
      
      <div class="tab-content" style="margin-top:1px;">
          <div class="tab-pane fade in active" id="clog">
              <pre id="log" class="pre-scrollable hide"></pre>
          </div>
        <div class="tab-pane fade in" id="cresult">
            <div id="download" class="col-sm-offset-1 hide">
                <p class="text-danger">下载链接:</p>
            </div>
        </div>
      </div>
    </div>
  </div>

  <footer class="navbar-fixed-bottom" style="
    line-height: 1.8;
    text-align: center;
    padding: 50px 0;
    color: #999;z-index: -99999; position:inherit;">
    <div class="author text-center">
      Official website:
      <a href="http://{{.Website}}">{{.Website}}</a> /
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>
</div>
  <!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
  <script src="https://cdn.bootcss.com/jquery/3.4.1/jquery.min.js"></script>
  <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
  <script src="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
<script src="/static/js/eshop.js"></script>
</body>
</html>
