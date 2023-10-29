package gamePhysics

import (
	"sync"
)

type Entity interface {
	Start()
	Update()
	SetID(ID int)
	GetID() int
	Finalize()
}

var entities []Entity
var mu sync.Mutex

func GetEntities() *[]Entity {
	if entities == nil {
		mu.Lock()
		if entities != nil {
			return &entities
		}
		entities = make([]Entity, 0)
		mu.Unlock()
	}

	return &entities
}

func AppendEntity(entity Entity) {
	mu.Lock()
	defer mu.Unlock()

	entities = append(entities, entity)
	entity.SetID(len(entities) - 1)
}

func DestroyEntity(entity Entity) {
	mu.Lock()
	defer mu.Unlock()

	entity.Finalize()

	freeIndex := entity.GetID()
	lastIndex := len(entities) - 1
	entities[lastIndex].SetID(freeIndex)

	entities[freeIndex] = entities[lastIndex]
	entities[lastIndex] = nil
	entities = entities[:lastIndex]
}
