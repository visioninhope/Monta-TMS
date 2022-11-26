# Generated by Django 4.1.3 on 2022-11-26 02:54

from django.db import migrations
import utils.models


class Migration(migrations.Migration):

    dependencies = [
        ("movements", "0001_initial"),
    ]

    operations = [
        migrations.AlterField(
            model_name="movement",
            name="status",
            field=utils.models.ChoiceField(
                choices=[
                    ("N", "New"),
                    ("P", "In Progress"),
                    ("C", "Completed"),
                    ("H", "Hold"),
                    ("B", "Billed"),
                    ("V", "Voided"),
                ],
                default="N",
                help_text="Status of the Movement",
                max_length=1,
                verbose_name="Status",
            ),
        ),
    ]