# Generated by Django 2.1.2 on 2019-08-05 03:27

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('cloud_provider', '0013_remove_plan_deploy_template'),
        ('kubeops_api', '0008_auto_20190724_0800'),
    ]

    operations = [
        migrations.AddField(
            model_name='cluster',
            name='plan',
            field=models.ForeignKey(null=True, on_delete=django.db.models.deletion.SET_NULL, to='cloud_provider.Plan'),
        ),
        migrations.AddField(
            model_name='cluster',
            name='worker_size',
            field=models.IntegerField(default=3),
        ),
    ]
