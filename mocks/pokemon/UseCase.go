// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo"
	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// CreatePokemon provides a mock function with given fields: _a0
func (_m *UseCase) CreatePokemon(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePokemon provides a mock function with given fields: _a0
func (_m *UseCase) DeletePokemon(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllPokemons provides a mock function with given fields: _a0
func (_m *UseCase) GetAllPokemons(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPokemon provides a mock function with given fields: _a0
func (_m *UseCase) GetPokemon(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePokemon provides a mock function with given fields: _a0
func (_m *UseCase) UpdatePokemon(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
