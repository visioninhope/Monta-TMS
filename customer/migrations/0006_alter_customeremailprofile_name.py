# Generated by Django 4.1.4 on 2022-12-16 17:57

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("customer", "0005_alter_customeremailprofile_name"),
    ]

    operations = [
        migrations.AlterField(
            model_name="customeremailprofile",
            name="name",
            field=models.CharField(
                help_text="Name", max_length=50, verbose_name="Name"
            ),
        ),
    ]
