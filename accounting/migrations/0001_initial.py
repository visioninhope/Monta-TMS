# Generated by Django 4.1.3 on 2022-12-04 00:03

import uuid

import django.core.validators
import django.db.models.deletion
import django_extensions.db.fields
from django.db import migrations, models

import utils.models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("organization", "0001_initial"),
    ]

    operations = [
        migrations.CreateModel(
            name="GeneralLedgerAccount",
            fields=[
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "id",
                    models.UUIDField(
                        default=uuid.uuid4,
                        editable=False,
                        primary_key=True,
                        serialize=False,
                        unique=True,
                    ),
                ),
                (
                    "is_active",
                    models.BooleanField(
                        default=True,
                        help_text="Designates whether this account is active.",
                        verbose_name="Active",
                    ),
                ),
                (
                    "account_number",
                    models.CharField(
                        help_text="The account number of the general ledger account.",
                        max_length=20,
                        unique=True,
                        validators=[
                            django.core.validators.RegexValidator(
                                message="Account number must be in the format 0000-0000-0000-0000.",
                                regex="^[0-9]{4}-[0-9]{4}-[0-9]{4}-[0-9]{4}$",
                            )
                        ],
                        verbose_name="Account Number",
                    ),
                ),
                (
                    "description",
                    models.CharField(
                        help_text="The description of the general ledger account.",
                        max_length=100,
                        verbose_name="Description",
                    ),
                ),
                (
                    "account_type",
                    utils.models.ChoiceField(
                        choices=[
                            ("ASSET", "Asset"),
                            ("LIABILITY", "Liability"),
                            ("EQUITY", "Equity"),
                            ("REVENUE", "Revenue"),
                            ("EXPENSE", "Expense"),
                        ],
                        help_text="The type of the general ledger account.",
                        max_length=9,
                        verbose_name="Account Type",
                    ),
                ),
                (
                    "cash_flow_type",
                    utils.models.ChoiceField(
                        blank=True,
                        choices=[
                            ("OPERATING", "Operating"),
                            ("INVESTING", "Investing"),
                            ("FINANCING", "Financing"),
                        ],
                        help_text="The cash flow type of the general ledger account.",
                        max_length=9,
                        verbose_name="Cash Flow Type",
                    ),
                ),
                (
                    "account_sub_type",
                    utils.models.ChoiceField(
                        blank=True,
                        choices=[
                            ("CURRENT_ASSET", "Current Asset"),
                            ("FIXED_ASSET", "Fixed Asset"),
                            ("OTHER_ASSET", "Other Asset"),
                            ("CURRENT_LIABILITY", "Current Liability"),
                            ("LONG_TERM_LIABILITY", "Long Term Liability"),
                            ("EQUITY", "Equity"),
                            ("REVENUE", "Revenue"),
                            ("COST_OF_GOODS_SOLD", "Cost of Goods Sold"),
                            ("EXPENSE", "Expense"),
                            ("OTHER_INCOME", "Other Income"),
                            ("OTHER_EXPENSE", "Other Expense"),
                        ],
                        help_text="The sub type of the general ledger account.",
                        max_length=19,
                        verbose_name="Account Sub Type",
                    ),
                ),
                (
                    "account_classification",
                    utils.models.ChoiceField(
                        blank=True,
                        choices=[
                            ("BANK", "Bank"),
                            ("CASH", "Cash"),
                            ("ACCOUNTS_RECEIVABLE", "Accounts Receivable"),
                            ("ACCOUNTS_PAYABLE", "Accounts Payable"),
                            ("INVENTORY", "Inventory"),
                            ("OTHER_CURRENT_ASSET", "Other Current Asset"),
                            ("FIXED_ASSET", "Fixed Asset"),
                        ],
                        help_text="The classification of the general ledger account.",
                        max_length=19,
                        verbose_name="Account Classification",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "General Ledger Account",
                "verbose_name_plural": "General Ledger Accounts",
                "ordering": ["account_number"],
            },
        ),
        migrations.CreateModel(
            name="RevenueCode",
            fields=[
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "id",
                    models.UUIDField(
                        default=uuid.uuid4,
                        editable=False,
                        primary_key=True,
                        serialize=False,
                        unique=True,
                    ),
                ),
                (
                    "code",
                    models.CharField(
                        help_text="The revenue code.",
                        max_length=4,
                        unique=True,
                        verbose_name="Code",
                    ),
                ),
                (
                    "description",
                    models.CharField(
                        help_text="The description of the revenue code.",
                        max_length=100,
                        verbose_name="Description",
                    ),
                ),
                (
                    "expense_account",
                    models.ForeignKey(
                        help_text="The expense account associated with the revenue code.",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="revenue_code_expense_account",
                        related_query_name="revenue_code_expense_accounts",
                        to="accounting.generalledgeraccount",
                        verbose_name="Expense Account",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
                (
                    "revenue_account",
                    models.ForeignKey(
                        help_text="The revenue account associated with the revenue code.",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="revenue_code_revenue_account",
                        related_query_name="revenue_code_revenue_accounts",
                        to="accounting.generalledgeraccount",
                        verbose_name="Revenue Account",
                    ),
                ),
            ],
            options={
                "verbose_name": "Revenue Code",
                "verbose_name_plural": "Revenue Codes",
                "ordering": ["code"],
            },
        ),
    ]
