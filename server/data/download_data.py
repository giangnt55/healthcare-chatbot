from datasets import load_dataset

dataset = load_dataset("wangrongsheng/HealthCareMagic-100k-en")
# In ra thông tin của dataset
print(dataset)

# Lưu dữ liệu thành các file (JSON hoặc CSV) nếu cần
dataset['train'].to_json('healthcare_magic_train.json')
dataset['test'].to_json('healthcare_magic_test.json')