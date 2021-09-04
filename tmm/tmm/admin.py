from django.contrib import admin
from django.apps import apps

class MyAdminSite(admin.AdminSite):

    def get_app_list(self, request):
        app_dict = self._build_app_dict(request)
        for k, v in app_dict.items():
            app = apps.get_app_config(v['app_label'])
            main_menu_index = getattr(app, 'main_menu_index', 9999)
            app_dict[k]['main_menu_index'] = main_menu_index
        app_list = sorted(app_dict.values(), key=lambda x: x['main_menu_index'])
        return app_list

    def index(self, request, extra_context=None):
        if extra_context is None:
            extra_context = {'home':'/admin_index'}

        app_list = self.get_app_list(request)
        
        #将数据写入到里面来到首页数据里
        xapps = []
        for app in app_list:
            tmp = []
            for m in app.get('models', []):
                model = apps.get_model(app_label=app.get('app_label'), model_name=m.get('object_name'))
                #查看是否有图
                if hasattr(model, 'show_plots'):
                    m['plots'] = model.show_plots()
                #是否有数值
                if hasattr(model, 'show_count'):
                    m['count'] = model.show_count()
                tmp.append(m)
            app['models'] = tmp 
            xapps.append(app)
        
        app_dict = self._build_app_dict(request)
        extra_context['app_list'] = xapps
        return super(MyAdminSite, self).index(request, extra_context)