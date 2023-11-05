package entities

import "sync"

var asteroids []*Asteroid
var mu sync.Mutex

func GetCopyAsteroids() []*Asteroid {
	if asteroids == nil {
		mu.Lock()
		if asteroids != nil {
			copyAsteroids := make([]*Asteroid, len(asteroids))
			copy(copyAsteroids, asteroids)
			return copyAsteroids
		}

		asteroids = make([]*Asteroid, 0)
		mu.Unlock()
	}

	copyAsteroids := make([]*Asteroid, len(asteroids))
	copy(copyAsteroids, asteroids)
	return copyAsteroids
}

func AppendAsteroid(a *Asteroid) {
	mu.Lock()
	defer mu.Unlock()

	asteroids = append(asteroids, a)
	a.SetID(len(asteroids) - 1)
}

func DestroyAsteroid(a Asteroid) {
	mu.Lock()
	defer mu.Unlock()

	a.Finalize()

	freeIndex := a.GetID()
	lastIndex := len(asteroids) - 1
	asteroids[lastIndex].SetID(freeIndex)

	asteroids[freeIndex] = asteroids[lastIndex]
	//asteroids[lastIndex] = nil
	asteroids = asteroids[:lastIndex]
}
