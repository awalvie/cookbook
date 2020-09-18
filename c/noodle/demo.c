#include <stdbool.h>
#include <SDL2/SDL.h>
#include <stdio.h>

#define SCREEN_WIDTH 640
#define SCREEN_HEIGHT 480

SDL_Window *window = NULL;
SDL_Surface *surface = NULL;
SDL_Surface *image = NULL;

int error(char *msg, const char *err)
{
	printf("Error %s: %s\n", msg, err);
}

bool init()
{
	bool success = true;
	if (SDL_Init(SDL_INIT_VIDEO) < 0) {
		error("init", SDL_GetError());
		success = false;
	} else {
		window = SDL_CreateWindow("Blank Window",
					  SDL_WINDOWPOS_UNDEFINED,
					  SDL_WINDOWPOS_UNDEFINED, SCREEN_WIDTH,
					  SCREEN_HEIGHT, SDL_WINDOW_SHOWN);

		if (window == NULL) {
			error("window", SDL_GetError());
			success = false;

		} else {
			surface = SDL_GetWindowSurface(window);
		}
	}
}

bool loadMedia()
{
	bool success = true;
	image = SDL_LoadBMP("graphic.bmp");
	if (image == NULL) {
		error("image", SDL_GetError());
		success = false;
	}
}

bool close()
{
	SDL_FreeSurface(surface);
	surface = NULL;
	SDL_DestroyWindow(window);
	window = NULL;
	SDL_Quit();
}
int main()
{
	if (!init()) {
		fprintf(stderr, "ERROR: Could not initialize SDL");
	} else {
		if (!loadMedia()) {
			fprintf(stderr, "ERROR: Could not load image");
		} else {
			SDL_BlitSurface(image, NULL, surface, NULL);
			SDL_UpdateWindowSurface(window);
			SDL_Delay(2000);
		}
	}

	close();
	return 0;
}
