try:
  import simplejson as json
except Exception as e:
  import json
from datetime import datetime as dt
import requests
from django.contrib import admin
from django.conf import settings
from django import forms
from .models import Online
#from flyadmin.widget.forms import SelectBoxWidget, TimelineWidget, EditorWidget, DateTimeWidget, UploadImagesWidget, InputNumberWidget, UploadFileWidget, StepsWidget, StepsNormalWidget
 

class OnlineAdminForm(forms.ModelForm):
  class Meta:
      model = Online
      widgets = {} 
      fields = '__all__'
  
  def __init__(self, *args, **kwargs):
      super(OnlineAdminForm, self).__init__(*args, **kwargs)
  
class OnlineAdmin(admin.ModelAdmin):
    form = OnlineAdminForm 
    list_display = ('name', 'server', 'cpu', 'memery', 'disk', 'disk2', 'status', 'update_time')
    search_fields = ('name',)
    def get_form(self, request, obj=None, **kwargs):
       form = super(OnlineAdmin, self).get_form(request, obj, **kwargs)
       return form



admin.site.register(Online, OnlineAdmin)
