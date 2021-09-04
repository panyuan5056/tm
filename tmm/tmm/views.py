from django.shortcuts import render
from django.apps import apps


 
def admin_index(request):
    #将数据写入到里面来到首页数据里(不得修改)
    html = ['<el-row>']
    for app in apps.get_app_configs():
        for model in app.get_models():
            if hasattr(model, 'show_plots'):
                show_plots = model.show_plots()
                if show_plots:
                    for plot in show_plots:
                        html.append('<div class="el-col el-col-{}">{}</div>'.format(plot.get('size'), plot.get('plot')))
    html.append('</el-row>')
    xplots = ''.join(html)   
    return render(request, "admin/xplots.html", locals()) 
 