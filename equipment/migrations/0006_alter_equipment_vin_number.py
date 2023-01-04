# Generated by Django 4.1.4 on 2023-01-03 16:54

from django.db import migrations, models

import equipment.validators


class Migration(migrations.Migration):

    dependencies = [
        ("equipment", "0005_alter_equipmenttypedetail_equipment_class"),
    ]

    operations = [
        migrations.AlterField(
            model_name="equipment",
            name="vin_number",
            field=models.CharField(
                blank=True,
                help_text="VIN number of the equipment.",
                max_length=17,
                validators=[equipment.validators.us_vin_number_validator],
                verbose_name="VIN Number",
            ),
        ),
    ]