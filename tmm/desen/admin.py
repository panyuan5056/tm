import json
from datetime import datetime as dt
import requests
from django.contrib import admin
from django.conf import settings
from django import forms
from .models import  Dense



class DenseAdminForm(forms.ModelForm):
  class Meta:
      model = Dense
      widgets = {} 
      fields = '__all__'
  
  def __init__(self, *args, **kwargs):
      super(DenseAdminForm, self).__init__(*args, **kwargs)
  
class DenseAdmin(admin.ModelAdmin):
    form = DenseAdminForm 
    list_display = ('title', 'create_time')
    search_fields = ('title',)

admin.site.register(Dense, DenseAdmin)

 
