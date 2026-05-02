package redis

type RedisRepo struct {
	//todo menambahkan redis client
}

func (r *RedisRepo) GetPricingFromCache(key string) (float64, error) {
	return 0.0, nil
}
