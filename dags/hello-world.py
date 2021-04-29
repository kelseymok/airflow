from airflow import DAG
from airflow.operators.bash_operator import BashOperator
from datetime import datetime, timedelta


default_args = {
  'owner': 'airflow',
  'depends_on_past': False,
  'start_date': datetime(2018, 8, 1),
  'email': ['foo@example.com'],
  'email_on_failure': False,
  'email_on_retry': False,
  'retries': 1,
  'retry_delay': timedelta(minutes=5),
}

dag = DAG('helloworld', default_args=default_args)

templated_command = """
    echo "Hello World"
"""

t1 = BashOperator(
    task_id='echo_hello_world',
    bash_command="/usr/local/airflow/src/hello-world.sh Zebra",
    dag=dag)

