#include <stdbool.h>
#include <SDL2/SDL.h>
#include <stdio.h>

static int SCREEN_WIDTH = 640;
static int SCREEN_HEIGHT = 480;
uint32_t *pixels;
bool gQuit = false;
bool leftMouseButtonDown = false;

SDL_Window *gWindow = NULL;
SDL_Renderer *gRenderer = NULL;
SDL_Texture *gTexture = NULL;

void error(char *msg, const char *err)
{
	printf("Error %s: %s\n", msg, err);
}

int init_array()
{
	int i, j;
	pixels = (uint32_t *)malloc(SCREEN_WIDTH * SCREEN_HEIGHT *
				    sizeof(uint32_t));

	if (pixels == NULL) {
		error("init_array", "could not allocate memory for pixels");
		return false;
	}

	for (i = 0; i < SCREEN_HEIGHT; i++) {
		for (j = 0; j < SCREEN_WIDTH; j++) {
			pixels[i * SCREEN_WIDTH + j] = 0xffffff;
		}
	}

	return true;
}

bool init()
{
	if (SDL_Init(SDL_INIT_VIDEO) < 0) {
		error("init", SDL_GetError());
		return false;
	}

	gWindow = SDL_CreateWindow("noodle", SDL_WINDOWPOS_UNDEFINED,
				   SDL_WINDOWPOS_UNDEFINED, SCREEN_WIDTH,
				   SCREEN_HEIGHT, SDL_WINDOW_SHOWN);

	if (gWindow == NULL) {
		error("window", SDL_GetError());
		return false;
	}

	gRenderer = SDL_CreateRenderer(gWindow, -1, 0);

	if (gRenderer == NULL) {
		error("renderer", SDL_GetError());
		return false;
	}

	gTexture = SDL_CreateTexture(gRenderer, SDL_PIXELFORMAT_ARGB8888,
				     SDL_TEXTUREACCESS_STATIC, SCREEN_WIDTH,
				     SCREEN_HEIGHT);

	if (gTexture == NULL) {
		error("texture", SDL_GetError());
		return false;
	}

	init_array();

	return true;
}

void quit()
{
	free(pixels);
	SDL_DestroyTexture(gTexture);
	gTexture = NULL;
	SDL_DestroyRenderer(gRenderer);
	gRenderer = NULL;
	SDL_DestroyWindow(gWindow);
	gWindow = NULL;
	SDL_Quit();
}

void end_loop()
{
	gQuit = true;
}

void handle_mouse(SDL_Event *event)
{
	int mouseX, mouseY;

	switch (event->type) {
	case SDL_QUIT:
		gQuit = true;
		break;
	case SDL_MOUSEBUTTONUP:
		if (event->button.button == SDL_BUTTON_LEFT)
			leftMouseButtonDown = false;
		break;
	case SDL_MOUSEBUTTONDOWN:
		if (event->button.button == SDL_BUTTON_LEFT)
			leftMouseButtonDown = true;
	case SDL_MOUSEMOTION:
		if (leftMouseButtonDown) {
			int mouseX = event->motion.x;
			int mouseY = event->motion.y;
			pixels[mouseY * SCREEN_WIDTH + mouseX] = 0;
		}
		break;
	}
}

void handle_keypress(SDL_Event *event)
{
	switch (event->key.keysym.sym) {
	case SDLK_q:
		end_loop();
		break;
	}
}

int main()
{
	if (!init()) {
		fprintf(stderr, "ERROR: Could not initialize SDL\n");
		quit();
		return 0;
	}
	/* main game loop */
	SDL_Event event;
	while (!gQuit) {
		SDL_UpdateTexture(gTexture, NULL, pixels,
				  SCREEN_WIDTH * sizeof(uint32_t));
		while (SDL_PollEvent(&event) != 0) {
			if (event.type == SDL_MOUSEBUTTONUP ||
			    event.type == SDL_MOUSEBUTTONDOWN ||
			    event.type == SDL_MOUSEMOTION) {
				handle_mouse(&event);
			} else if (event.type == SDL_KEYDOWN) {
				handle_keypress(&event);
			}
		}
		SDL_RenderClear(gRenderer);
		SDL_RenderCopy(gRenderer, gTexture, NULL, NULL);
		SDL_RenderPresent(gRenderer);
	}

	quit();
	return 0;
}
