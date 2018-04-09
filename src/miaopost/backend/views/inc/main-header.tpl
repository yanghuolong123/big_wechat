<!-- Main Header -->
  <header class="main-header">

    {{template "inc/logo.tpl" .}}

    <!-- Header Navbar -->
    <nav class="navbar navbar-static-top" role="navigation">
    
      <!-- Sidebar toggle button-->
      <a href="#" class="sidebar-toggle" data-toggle="offcanvas" role="button">
        <span class="sr-only">Toggle navigation</span>
      </a>
      
      <!-- Navbar Right Menu -->
      <div class="navbar-custom-menu">
        <ul class="nav navbar-nav">               
          
          <!-- User Account Menu -->
          <li class="dropdown user user-menu">
            <!-- Menu Toggle Button -->
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">              
              <img src="/static/plugin/adminlte/dist/img/user2-160x160.jpg" class="user-image" alt="User Image">              
              <span class="hidden-xs">{{.user.Username}}</span>
            </a>
            <ul class="dropdown-menu">
              <!-- The user image in the menu -->
              <li class="user-header">
                <img src="/static/plugin/adminlte/dist/img/user2-160x160.jpg" class="img-circle" alt="User Image">
                <p>
                 {{if .user}} 
                  <small>登陆时间: {{date .user.Last_logintime "Y-m-d H:i:s"}}</small>
		 {{end}}
                </p>
              </li>
              
              <!-- Menu Footer-->
              <li class="user-footer">
                <div class="pull-left">
                  <a href="/" class="btn btn-default btn-flat">修改密码</a>
                </div>
                <div class="pull-right">
                  <a href="/logout" class="btn btn-default btn-flat">退出</a>
                </div>
              </li>
            </ul>
          </li>
          
        </ul>
      </div>
    </nav>
  </header>
