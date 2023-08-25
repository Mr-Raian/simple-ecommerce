# Simple Ecommerce API

An order placing and tracking API that offers essential functionality. This is running at production in [freedomof.tech](https://freedomof.tech/)

This project features Unit and Integration testing, and was designed with SOLID principles in mind.

#### Endpoints are:
- `/contactForm`: Submit contact forms.
- `/items/get_data`: Retrieve item data.
- `/orders/new`: Place new orders.
- `/orders/info`: Retrieve order information.
- `/health`: Check system health.
- `/stripeWebhook`: Handle Stripe webhook events.

#### Main Dependencies:
- PostgreSQL: as the database.
- [Echo](https://echo.labstack.com/): as the http framework.
- [Amazon Simple Email Service (SES)](https://aws.amazon.com/ses/): for sending emails.
