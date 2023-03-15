# Generated by Django 4.1.7 on 2023-03-11 09:23

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("stops", "0001_initial"),
    ]

    operations = [
        migrations.AlterField(
            model_name="qualifiercode",
            name="code",
            field=models.CharField(
                help_text="Code of the Qualifier Code",
                max_length=255,
                verbose_name="Code",
            ),
        ),
        migrations.AddConstraint(
            model_name="qualifiercode",
            constraint=models.UniqueConstraint(
                fields=("code", "organization"),
                name="unique_qualifier_code_organization",
            ),
        ),
    ]