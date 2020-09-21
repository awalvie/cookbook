#include <stdbool.h>
#include <SDL2/SDL.h>
#include <stdio.h>

#define SCREEN_WIDTH 640
#define SCREEN_HEIGHT 480

SDL_Window *gWindow = NULL;
SDL_Surface *gSurface = NULL;
SDL_Surface *gCanvas = NULL;

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
		gWindow =
			SDL_CreateWindow("noodle", SDL_WINDOWPOS_UNDEFINED,
					 SDL_WINDOWPOS_UNDEFINED, SCREEN_WIDTH,
					 SCREEN_HEIGHT, SDL_WINDOW_SHOWN);

		if (gWindow == NULL) {
			error("window", SDL_GetError());
			success = false;
		} else {
			gSurface = SDL_GetWindowSurface(gWindow);
			if (gSurface == NULL) {
				error("surface", SDL_GetError());
				success = false;
			}
		}
	}

	return success;
}

bool loadMedia()
{
	bool success = true;

	gCanvas = SDL_CreateRGBSurface(0, SCREEN_WIDTH, SCREEN_HEIGHT, 32, 0, 0,
				       0, 0);
	SDL_FillRect(gCanvas, NULL,
		     SDL_MapRGB(gCanvas->format, 0xFF, 0xFF, 0xFF));

	if (gCanvas == NULL) {
		error("canvas", SDL_GetError());
		success = false;
	}

	return success;
}

bool quit()
{
	SDL_FreeSurface(gCanvas);
	gCanvas = NULL;
	SDL_DestroyWindow(gWindow);
	gWindow = NULL;
	SDL_Quit();
}

int main()
{
	if (!init()) {
		fprintf(stderr, "ERROR: Could not initialize SDL\n");
	} else {
		if (!loadMedia()) {
			fprintf(stderr, "ERROR: Could not load image\n");
		} else {
			SDL_BlitSurface(gCanvas, NULL, gSurface, NULL);
			SDL_UpdateWindowSurface(gWindow);
			/* main game loop */
			bool quit = false;
			SDL_Event e;
			while (!quit) {
				while (SDL_PollEvent(&e) != 0) {
					if (e.type == SDL_MOUSEMOTION ||
					    e.type == SDL_MOUSEBUTTONDOWN ||
					    e.type == SDL_MOUSEBUTTONUP) {
						int x, y;
						SDL_GetMouseState(&x, &y);
						printf("X: %d\nY: %d\n", x, y);
					}
					else if (e.type == SDL_KEYDOWN) {
						if (e.key.keysym.sym == SDLK_q) {
							quit = true;
						}
					}
				}
			}
		}
	}

	quit();
	return 0;
}
