from datasets import load_dataset

dataset = load_dataset("wangrongsheng/HealthCareMagic-100k-en")
dataset['train'].to_json('healthcare_magic_train.json')
dataset['test'].to_json('healthcare_magic_test.json')