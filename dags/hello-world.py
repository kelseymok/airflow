from airflow import DAG
from airflow.operators.bash_operator import BashOperator
from datetime import datetime, timedelta

# Use the best practices for python -> classes and call classes
default_args = {
  'owner': 'airflow',
  'depends_on_past': False,
  'start_date': datetime(2020, 4, 29),
  'email': ['foo@example.com'],
  'email_on_failure': False,
  'email_on_retry': False,
  'retries': 1,
  'retry_delay': timedelta(minutes=5),
}

dag = DAG('helloworld', default_args=default_args)


t1 = BashOperator(
    task_id='zebra',
    bash_command="/opt/airflow/plugins/hello-world.sh Zebra",
    dag=dag)


t2 = BashOperator(
    task_id='horse',
    bash_command="/opt/airflow/plugins/hello-world.sh Horse",
    dag=dag)

t1 >> t2