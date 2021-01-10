package favpokemon

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ricardoham/pokedex-api/entity"
	repository "github.com/ricardoham/pokedex-api/infrastructure/repository"
)

type FavPokemonService struct {
	repository *repository.FavPokemonsRepository
}

func NewFavPokemonsService(repository *repository.FavPokemonsRepository) *FavPokemonService {
	return &FavPokemonService{
		repository: repository,
	}
}

func (s *FavPokemonService) CreateFavPokemon(pokemon *entity.FavPokemon) error {
	p, err := entity.NewFavPokemon(pokemon.Name, time.Now())
	if err != nil {
		return err
	}

	return s.repository.Create(p)
}

func (s *FavPokemonService) GetAllFavPokemons() ([]*entity.FavPokemon, error) {
	var pokemons []*entity.FavPokemon
	ctx := context.TODO()
	pokemons, err := s.repository.FindAll(ctx, pokemons)

	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (s *FavPokemonService) UpdateFavPokemon(pokeId uuid.UUID, updateData *entity.FavPokemon) (int64, error) {
	ctx := context.TODO()
	result, err := s.repository.Update(ctx, pokeId, updateData)
	if err != nil {
		return 0, err
	}

	return result.MatchedCount, nil
}

func (s *FavPokemonService) DeleteFavPokemon(pokeId entity.DeleteFavPokemon) (int64, error) {
	ctx := context.TODO()
	deleteResult, err := s.repository.Delete(ctx, pokeId)
	if err != nil {
		return 0, err
	}

	return deleteResult.DeletedCount, nil
}
