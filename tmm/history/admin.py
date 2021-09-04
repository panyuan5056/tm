import json
from datetime import datetime as dt
import requests
from django.contrib import admin
from django.conf import settings
from django import forms
from .models import  History



class HistoryAdminForm(forms.ModelForm):
  class Meta:
      model = History
      widgets = {} 
      fields = '__all__'
  
  def __init__(self, *args, **kwargs):
      super(HistoryAdminForm, self).__init__(*args, **kwargs)
  
class HistoryAdmin(admin.ModelAdmin):
    form = HistoryAdminForm 
    list_display = ('title', 'file', 'field', 'category', 'create_time')
    search_fields = ('title','file',)

admin.site.register(History, HistoryAdmin)



