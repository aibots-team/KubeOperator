# Generated by Django 2.1.2 on 2019-10-16 08:17

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('kubeops_api', '0031_merge_20191016_0817'),
    ]

    operations = [
        migrations.AlterModelOptions(
            name='node',
            options={'ordering': ['name']},
        ),
        migrations.AlterField(
            model_name='clusterhealthhistory',
            name='date_type',
            field=models.CharField(choices=[('HOUR', 'HOUR'), ('DAY', 'DAY')], default='HOUR', max_length=255),
        ),
    ]
